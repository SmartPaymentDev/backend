package handler

import "smart-payment-dev-be/internal/adapter/inbound/registry"

type Handler struct {
	serviceRegistry registry.ServiceRegistry
}

func New(serviceRegistry registry.ServiceRegistry) *Handler {
	return &Handler{
		serviceRegistry: serviceRegistry,
	}
}

func (h *Handler) GetServiceRegistry() registry.ServiceRegistry {
	return h.serviceRegistry
}
