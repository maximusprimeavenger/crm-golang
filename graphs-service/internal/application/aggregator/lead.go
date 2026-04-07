package aggregator

import (
	"fmt"
	domain "graphs-service/internal/entities"
	"graphs-service/internal/interfaces/mapper"
)

func (a *aggregator) eventLeadCreated(event domain.Event) {
	leadPayload, err := mapper.LeadDTOToDomain(event.Payload)
	if err != nil {
		a.logger.Error(fmt.Sprintf("error mapping item with id: %d", leadPayload.ID), "error", err)
	}
	_, ok := a.Leads[leadPayload.ID]
	if !ok {

	}
}
