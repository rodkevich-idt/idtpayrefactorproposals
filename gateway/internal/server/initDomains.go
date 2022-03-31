package server

import (
	antifraudHandler "gateway/internal/domain/antifraud/handler/http"
	antifraudRepo "gateway/internal/domain/antifraud/datasource/database"
	antifraudUseCase "gateway/internal/domain/antifraud/usecase"
	cardHandler "gateway/internal/domain/card/handler/http"
	cardRepo "gateway/internal/domain/card/datasource/database"
	cardUseCase "gateway/internal/domain/card/usecase"

	...
	...
	...
	...

	"net/http"

	"gateway/internal/middleware"
	"gateway/internal/utility/respond"
)

func (s *Server) InitDomains() {
	s.initSettlement()
	s.initMicrodeposit()
	s.initCardinfo()
	s.initAntifraud()
	s.initMerchant()
	s.initCardprovider()
	s.initCard()
	s.initCardprocessor()
	s.initWallet()
	s.initWalletprovider()
	s.initWalletprocessor()
	s.initToken()
	s.initUser()
	s.initVersion()
	s.initHealth()
}

func (s *Server) initVersion() {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			respond.Json(w, http.StatusOK, map[string]string{"version": s.version})
		})
	})
}

func (s *Server) initCardprocessor() {
	newCardprocessorRepo := cardprocessorRepo.New(s.DB())
	newCardprocessorUseCase := cardprocessorUseCase.New(newCardprocessorRepo)
	cardprocessorHandler.RegisterHTTPEndPoints(s.router, s.validator, newCardprocessorUseCase)
}

func (s *Server) initCardinfo() {
	newCardinfoRepo := cardinfoRepo.New(s.DB())
	newCardinfoUseCase := cardinfoUseCase.New(newCardinfoRepo)
	cardinfoHandler.RegisterHTTPEndPoints(s.router, s.validator, newCardinfoUseCase)
}

func (s *Server) initMicrodeposit() {
	newMicrodepositRepo := microdepositRepo.New(s.DB())
	newMicrodepositUseCase := microdepositUseCase.New(newMicrodepositRepo)
	microdepositHandler.RegisterHTTPEndPoints(s.router, s.validator, newMicrodepositUseCase)
}

func (s *Server) initSettlement() {
	newSettlementRepo := settlementRepo.New(s.DB())
	newSettlementUseCase := settlementUseCase.New(newSettlementRepo)
	settlementHandler.RegisterHTTPEndPoints(s.router, s.validator, newSettlementUseCase)
}
