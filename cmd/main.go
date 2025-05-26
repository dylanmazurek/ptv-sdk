package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
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
	fetchAndPrintRouteTypes(ptvClient)
	fetchAndPrintRoutes(ptvClient)
	fetchAndPrintDepartures(ptvClient)
}

func fetchAndPrintRouteTypes(client *ptv.Client) {
	routeTypes, err := client.RouteTypes()
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch route types")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Route Type Name", "Route Type"})
	for _, rt := range routeTypes {
		t.AppendRow(table.Row{rt.RouteTypeName, rt.RouteType})
	}

	t.Render()
}

func fetchAndPrintRoutes(client *ptv.Client) {
	reqParams := models.RoutesRequest{
		RouteTypes: []int{0, 1},
	}

	routes, err := client.Routes(reqParams)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch routes")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Route ID", "Route Name", "Route Type"})
	for _, route := range routes {
		t.AppendRow(table.Row{route.RouteID, route.RouteName, route.RouteType})
	}

	t.Render()
}

func fetchAndPrintDepartures(client *ptv.Client) {
	stopId := 2942
	routeId := 1041

	departures, err := client.DeparturesForStop(1, stopId, routeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch departures")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	fmt.Printf("Departures for Route %d at Stop %d\n", routeId, stopId)
	t.AppendHeader(table.Row{"Scheduled", "Estimated"})
	for _, dep := range departures {
		estimated := "-"
		if dep.EstimatedDeparture != nil {
			estimated = dep.EstimatedDeparture.Local().Format("03:04PM")
		}

		t.AppendRow(table.Row{dep.ScheduledDeparture.Local().Format("03:04PM"), estimated})
	}

	t.Render()
}
