package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/smithy-go/ptr"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models/filters"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).
		Level(zerolog.DebugLevel).
		With().
		Logger()

	timezone, err := time.LoadLocation("Australia/Melbourne")
	if err != nil {
		log.Error().Err(err).Msg("failed to load timezone, defaulting to UTC")
		timezone = time.UTC
	}

	ptvOpts := []ptv.Option{
		ptv.WithUserID(os.Getenv("PTV_USER_ID")),
		ptv.WithAccessKey(os.Getenv("PTV_API_KEY")),
		ptv.WithTimezone(timezone),
	}

	ptvClient := ptv.New(ctx, ptvOpts...)
	fetchAndPrintDepartures(ptvClient, timezone)
}

func fetchAndPrintDepartures(client *ptv.Client, timezone *time.Location) {
	stopID, _ := strconv.Atoi(os.Getenv("TEST_STOP_ID"))
	routeType, _ := strconv.Atoi(os.Getenv("TEST_ROUTE_TYPE_ID"))
	routeID, _ := strconv.Atoi(os.Getenv("TEST_ROUTE_ID"))

	afterDate := time.Now()

	f := &filters.DeparturesFilter{
		RouteType: constants.RouteType(routeType),
		StopID:    stopID,
		RouteID:   routeID,
		AfterDate: &afterDate,
		Expand: []constants.ExpandOption{
			constants.ExpandRoute,
			constants.ExpandStop,
			constants.ExpandDirection,
			constants.ExpandRun,
			constants.ExpandVehiclePosition,
		},

		BaseFilter: filters.BaseFilter{
			MaxResults: ptr.Int(5),
		},
	}

	resp, err := client.Departures(f)
	if err != nil {
		log.Error().Err(err).Msg("failed to perform search")
		return
	}

	if len(resp.Routes) == 0 {
		log.Warn().Msg("no routes found for the given stop")
		return
	}

	if len(resp.Stops) == 0 {
		log.Warn().Msg("no stop found for the given stop ID")
		return
	}

	route := resp.Routes[0]
	if len(resp.Routes) > 1 {
		log.Warn().Msg("multiple routes found for the same stop, displaying the first one")
	}

	stop := resp.Stops[0]
	if len(resp.Stops) > 1 {
		log.Warn().Msg("multiple stops found for the same stop ID, displaying the first one")
	}

	t1 := table.NewWriter()
	t1.SetOutputMirror(os.Stdout)
	t1.AppendHeader(table.Row{"", ""})
	t1.AppendRow(table.Row{"Route Type", route.RouteType.FriendlyString()})
	t1.AppendRow(table.Row{"Route", fmt.Sprintf("%s (%s)", route.Name, route.Number)})
	t1.AppendRow(table.Row{"Stop", stop.Name})

	t1.Render()

	t2 := table.NewWriter()
	t2.SetOutputMirror(os.Stdout)
	t2.AppendHeader(table.Row{"Departure", "Delayed"})
	for _, dep := range resp.Departures {
		delayMinStr := ""
		delayMin := dep.DelayMin()
		if delayMin != nil {
			delayMinStr = fmt.Sprintf("+%d min", *delayMin)
		}

		t2.AppendRow(table.Row{
			dep.FriendlyDepartureTime(timezone),
			delayMinStr,
		})
	}

	t2.Render()
}
