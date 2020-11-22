package types

import (
	"time"

	"github.com/twpayne/go-geom"
)

// A Section holds information about a specific section.
type Section struct {
	Type SectionType
	ID   ID
	Mode string

	// From & To
	From Container
	To   Container

	// Arrival time & departure time
	Departure time.Time
	Arrival   time.Time

	// Duration of travel
	Duration time.Duration

	// The path taken by this section
	Path []PathSegment

	// The path in geojson format
	Geo *geom.LineString

	// List of the stop times of this section
	StopTimes []StopTime

	// Information to display
	Display Display

	// Additional informations, from what I can see this is always a PTMethod
	Additional []PTMethod
}

// A SectionType codifies the type of section that can be encountered.
type SectionType string

// These are the types of sections that can be returned from the API.
const (
	// Public transport section.
	SectionPublicTransport SectionType = "public_transport"

	// Street section.
	SectionStreetNetwork SectionType = "street_network"

	// Waiting section between transport.
	SectionWaiting SectionType = "waiting"

	// This “stay in the vehicle” section occurs when the traveller has to stay in the vehicle when the bus change its routing.
	SectionStayIn SectionType = "stay_in"

	// Transfer section.
	SectionTransfer SectionType = "transfer"

	// Teleportation section. Used when starting or arriving to a city or a stoparea (“potato shaped” objects) Useful to make navitia idempotent.
	// Warning: Be careful: no Path nor Geo items in this case.
	SectionCrowFly SectionType = "crow_fly"

	// Vehicle may not drive along: traveler will have to call agency to confirm journey
	// Also sometimes called ODT.
	SectionOnDemandTransport SectionType = "on_demand_transport"

	// Taking a bike from a bike sharing system (bss).
	SectionBikeShareRent SectionType = "bss_rent"

	// Putting back a bike from a bike sharing system (bss).
	SectionBikeSharePutBack SectionType = "bss_put_back"

	// Boarding on plane.
	SectionBoarding SectionType = "boarding"

	// Landing off the plane.
	SectionLanding SectionType = "landing"
)

// SectionTypes is the type of a section.
var SectionTypes = map[SectionType]string{
	SectionPublicTransport:   "Public transport section",
	SectionStreetNetwork:     "Street section",
	SectionWaiting:           "Waiting section between transport",
	SectionStayIn:            "This “stay in the vehicle” section occurs when the traveller has to stay in the vehicle when the bus change its routing.",
	SectionTransfer:          "Transfer section",
	SectionCrowFly:           "Teleportation section. Used when starting or arriving to a city or a stoparea (“potato shaped” objects) Useful to make navitia idempotent",
	SectionOnDemandTransport: "Vehicle may not drive along: traveler will have to call agency to confirm journey",
	SectionBikeShareRent:     "Taking a bike from a bike sharing system (bss)",
	SectionBikeSharePutBack:  "Putting back a bike from a bike sharing system (bss)",
	SectionBoarding:          "Boarding on plane",
	SectionLanding:           "Landing off the plane",
}

// A StopTime stores info about a stop in a route: when the vehicle comes in, when it comes out, and what stop it is.
type StopTime struct {
	// The PTDateTime of the stop, this stores the info about the arrival & departure
	PTDateTime PTDateTime

	// The stop point in question
	StopPoint StopPoint `json:"stop_point"`

	UTCDepartureTime string `json:"utc_departure_time"`

	Headsign string `json:"headsign"`

	UTCArrivalTime string `json:"utc_arrival_time"`

	DepartureTime string `json:"departure_time"`

	PickupAllowed bool `json:"pickup_allowed"`

	DropOffAllowed bool `json:"drop_off_allowed"`
}

// A PTMethod is a Public Transportation method: it can be regular, estimated times or ODT (on-demand transport).
type PTMethod string

// PTMethodXXX codes for known PTMethod.
const (
	// PTMethodRegular: No on-demand transport. Line does not contain any estimated stop times, nor zonal stop point location. No need to call too.
	PTMethodRegular PTMethod = "regular"

	// PTMethodDateTimeEstimated: No on-demand transport. However, line has at least one estimated date time.
	PTMethodDateTimeEstimated PTMethod = "had_date_time_estimated"

	// PTMethodODTStopTime: Line does not contain any estimated stop times, nor zonal stop point location. But you will have to call to take it.
	PTMethodODTStopTime PTMethod = "odt_with_stop_time"

	// PTMethodODTStopPoint: Line can contain some estimated stop times, but no zonal stop point location. And you will have to call to take it.
	PTMethodODTStopPoint PTMethod = "odt_with_stop_point"

	// PTMethodODTZone: Line can contain some estimated stop times, and zonal stop point location. And you will have to call to take it. Well, not really a public transport line, more a cab….
	PTMethodODTZone PTMethod = "odt_with_zone"
)
