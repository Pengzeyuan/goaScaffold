package serializer

import (
	gensimulation "boot/gen/simulation"
	"boot/model"
	"encoding/json"
)

func SimulationModel2GetDataResp(simulation *model.Simulation) *gensimulation.GetDataResp {

	data := gensimulation.GetDataResp{
		Key:            simulation.Key,
		Val:            simulation.Val,
		IsShowMock:     simulation.IsShowMock,
		OrderBy:        &simulation.OrderBy,
		OrderTimeScope: &simulation.OrderTimeScope,
	}
	err := json.Unmarshal([]byte(simulation.Val), &data)
	if err != nil {
		return &data
	}
	return &data
}
