// Code generated by goa v3.2.4, DO NOT EDIT.
//
// entity_hall HTTP client CLI support package
//
// Command:
// $ goa gen boot/design

package client

import (
	entityhall "boot/gen/entity_hall"
	"encoding/json"
	"fmt"
)

// BuildWaitLineOverviewPayload builds the payload for the entity_hall
// WaitLineOverview endpoint from CLI flags.
func BuildWaitLineOverviewPayload(entityHallWaitLineOverviewBody string) (*entityhall.WaitLineOverviewPayload, error) {
	var err error
	var body WaitLineOverviewRequestBody
	{
		err = json.Unmarshal([]byte(entityHallWaitLineOverviewBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"endDate\": \"2020-07-27\",\n      \"regionCode\": \"520100\",\n      \"startDate\": \"2019-07-27\"\n   }'")
		}
	}
	v := &entityhall.WaitLineOverviewPayload{
		RegionCode: body.RegionCode,
		StartDate:  body.StartDate,
		EndDate:    body.EndDate,
	}

	return v, nil
}
