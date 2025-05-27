package main

import (
	"context"
	"os"
	"strconv"
	"time"

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

	ptvOpts := []ptv.Option{
		ptv.WithUserID(os.Getenv("PTV_USER_ID")),
		ptv.WithAccessKey(os.Getenv("PTV_API_KEY")),
		ptv.WithTimezone("Australia/Melbourne"),
	}

	ptvClient := ptv.New(ctx, ptvOpts...)
	fetchAndPrintDepartures(ptvClient)
}

func fetchAndPrintDepartures(client *ptv.Client) {
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
			constants.ExpandAll,
		},
	}

	resp, err := client.Departures(f)
	if err != nil {
		log.Error().Err(err).Msg("failed to perform search")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Direction ID", "Scheduled Departure", "Estimated Departure"})
	for _, dep := range resp.Departures {
		t.AppendRow(table.Row{
			dep.DirectionID,
			dep.ScheduledDeparture.HumanString(),
			dep.EstimatedDeparture.HumanString(),
		})
	}

	t.Render()
}
