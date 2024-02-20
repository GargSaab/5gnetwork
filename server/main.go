/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	AccessAndMobilitySubscriptionDataRetrievalAPIService := openapi.NewAccessAndMobilitySubscriptionDataRetrievalAPIService()
	AccessAndMobilitySubscriptionDataRetrievalAPIController := openapi.NewAccessAndMobilitySubscriptionDataRetrievalAPIController(AccessAndMobilitySubscriptionDataRetrievalAPIService)

	Class5MBSSubscriptionDataRetrievalAPIService := openapi.NewClass5MBSSubscriptionDataRetrievalAPIService()
	Class5MBSSubscriptionDataRetrievalAPIController := openapi.NewClass5MBSSubscriptionDataRetrievalAPIController(Class5MBSSubscriptionDataRetrievalAPIService)

	EnhancedCoverageRestrictionDataRetrievalAPIService := openapi.NewEnhancedCoverageRestrictionDataRetrievalAPIService()
	EnhancedCoverageRestrictionDataRetrievalAPIController := openapi.NewEnhancedCoverageRestrictionDataRetrievalAPIController(EnhancedCoverageRestrictionDataRetrievalAPIService)

	GPSIToSUPITranslationOrSUPIToGPSITranslationAPIService := openapi.NewGPSIToSUPITranslationOrSUPIToGPSITranslationAPIService()
	GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController := openapi.NewGPSIToSUPITranslationOrSUPIToGPSITranslationAPIController(GPSIToSUPITranslationOrSUPIToGPSITranslationAPIService)

	GroupIdentifiersAPIService := openapi.NewGroupIdentifiersAPIService()
	GroupIdentifiersAPIController := openapi.NewGroupIdentifiersAPIController(GroupIdentifiersAPIService)

	LCSBroadcastAssistanceDataTypesRetrievalAPIService := openapi.NewLCSBroadcastAssistanceDataTypesRetrievalAPIService()
	LCSBroadcastAssistanceDataTypesRetrievalAPIController := openapi.NewLCSBroadcastAssistanceDataTypesRetrievalAPIController(LCSBroadcastAssistanceDataTypesRetrievalAPIService)

	LCSMobileOriginatedDataRetrievalAPIService := openapi.NewLCSMobileOriginatedDataRetrievalAPIService()
	LCSMobileOriginatedDataRetrievalAPIController := openapi.NewLCSMobileOriginatedDataRetrievalAPIController(LCSMobileOriginatedDataRetrievalAPIService)

	LCSPrivacyDataRetrievalAPIService := openapi.NewLCSPrivacyDataRetrievalAPIService()
	LCSPrivacyDataRetrievalAPIController := openapi.NewLCSPrivacyDataRetrievalAPIController(LCSPrivacyDataRetrievalAPIService)

	MultipleIdentifiersAPIService := openapi.NewMultipleIdentifiersAPIService()
	MultipleIdentifiersAPIController := openapi.NewMultipleIdentifiersAPIController(MultipleIdentifiersAPIService)

	ProseSubscriptionDataRetrievalAPIService := openapi.NewProseSubscriptionDataRetrievalAPIService()
	ProseSubscriptionDataRetrievalAPIController := openapi.NewProseSubscriptionDataRetrievalAPIController(ProseSubscriptionDataRetrievalAPIService)

	ProvidingAcknowledgementOfCAGUpdateAPIService := openapi.NewProvidingAcknowledgementOfCAGUpdateAPIService()
	ProvidingAcknowledgementOfCAGUpdateAPIController := openapi.NewProvidingAcknowledgementOfCAGUpdateAPIController(ProvidingAcknowledgementOfCAGUpdateAPIService)

	ProvidingAcknowledgementOfSNSSAIsUpdateAPIService := openapi.NewProvidingAcknowledgementOfSNSSAIsUpdateAPIService()
	ProvidingAcknowledgementOfSNSSAIsUpdateAPIController := openapi.NewProvidingAcknowledgementOfSNSSAIsUpdateAPIController(ProvidingAcknowledgementOfSNSSAIsUpdateAPIService)

	ProvidingAcknowledgementOfSteeringOfRoamingAPIService := openapi.NewProvidingAcknowledgementOfSteeringOfRoamingAPIService()
	ProvidingAcknowledgementOfSteeringOfRoamingAPIController := openapi.NewProvidingAcknowledgementOfSteeringOfRoamingAPIController(ProvidingAcknowledgementOfSteeringOfRoamingAPIService)

	ProvidingAcknowledgementOfUEParametersUpdateAPIService := openapi.NewProvidingAcknowledgementOfUEParametersUpdateAPIService()
	ProvidingAcknowledgementOfUEParametersUpdateAPIController := openapi.NewProvidingAcknowledgementOfUEParametersUpdateAPIController(ProvidingAcknowledgementOfUEParametersUpdateAPIService)

	RetrievalOfMultipleDataSetsAPIService := openapi.NewRetrievalOfMultipleDataSetsAPIService()
	RetrievalOfMultipleDataSetsAPIController := openapi.NewRetrievalOfMultipleDataSetsAPIController(RetrievalOfMultipleDataSetsAPIService)

	RetrievalOfSharedDataAPIService := openapi.NewRetrievalOfSharedDataAPIService()
	RetrievalOfSharedDataAPIController := openapi.NewRetrievalOfSharedDataAPIController(RetrievalOfSharedDataAPIService)

	RetrievalOfTheIndividualSharedDataAPIService := openapi.NewRetrievalOfTheIndividualSharedDataAPIService()
	RetrievalOfTheIndividualSharedDataAPIController := openapi.NewRetrievalOfTheIndividualSharedDataAPIController(RetrievalOfTheIndividualSharedDataAPIService)

	SMFSelectionSubscriptionDataRetrievalAPIService := openapi.NewSMFSelectionSubscriptionDataRetrievalAPIService()
	SMFSelectionSubscriptionDataRetrievalAPIController := openapi.NewSMFSelectionSubscriptionDataRetrievalAPIController(SMFSelectionSubscriptionDataRetrievalAPIService)

	SMSManagementSubscriptionDataRetrievalAPIService := openapi.NewSMSManagementSubscriptionDataRetrievalAPIService()
	SMSManagementSubscriptionDataRetrievalAPIController := openapi.NewSMSManagementSubscriptionDataRetrievalAPIController(SMSManagementSubscriptionDataRetrievalAPIService)

	SMSSubscriptionDataRetrievalAPIService := openapi.NewSMSSubscriptionDataRetrievalAPIService()
	SMSSubscriptionDataRetrievalAPIController := openapi.NewSMSSubscriptionDataRetrievalAPIController(SMSSubscriptionDataRetrievalAPIService)

	SessionManagementSubscriptionDataRetrievalAPIService := openapi.NewSessionManagementSubscriptionDataRetrievalAPIService()
	SessionManagementSubscriptionDataRetrievalAPIController := openapi.NewSessionManagementSubscriptionDataRetrievalAPIController(SessionManagementSubscriptionDataRetrievalAPIService)

	SliceSelectionSubscriptionDataRetrievalAPIService := openapi.NewSliceSelectionSubscriptionDataRetrievalAPIService()
	SliceSelectionSubscriptionDataRetrievalAPIController := openapi.NewSliceSelectionSubscriptionDataRetrievalAPIController(SliceSelectionSubscriptionDataRetrievalAPIService)

	SubscriptionCreationAPIService := openapi.NewSubscriptionCreationAPIService()
	SubscriptionCreationAPIController := openapi.NewSubscriptionCreationAPIController(SubscriptionCreationAPIService)

	SubscriptionCreationForSharedDataAPIService := openapi.NewSubscriptionCreationForSharedDataAPIService()
	SubscriptionCreationForSharedDataAPIController := openapi.NewSubscriptionCreationForSharedDataAPIController(SubscriptionCreationForSharedDataAPIService)

	SubscriptionDeletionAPIService := openapi.NewSubscriptionDeletionAPIService()
	SubscriptionDeletionAPIController := openapi.NewSubscriptionDeletionAPIController(SubscriptionDeletionAPIService)

	SubscriptionDeletionForSharedDataAPIService := openapi.NewSubscriptionDeletionForSharedDataAPIService()
	SubscriptionDeletionForSharedDataAPIController := openapi.NewSubscriptionDeletionForSharedDataAPIController(SubscriptionDeletionForSharedDataAPIService)

	SubscriptionModificationAPIService := openapi.NewSubscriptionModificationAPIService()
	SubscriptionModificationAPIController := openapi.NewSubscriptionModificationAPIController(SubscriptionModificationAPIService)

	TraceConfigurationDataRetrievalAPIService := openapi.NewTraceConfigurationDataRetrievalAPIService()
	TraceConfigurationDataRetrievalAPIController := openapi.NewTraceConfigurationDataRetrievalAPIController(TraceConfigurationDataRetrievalAPIService)

	TriggerSORInfoUpdateAPIService := openapi.NewTriggerSORInfoUpdateAPIService()
	TriggerSORInfoUpdateAPIController := openapi.NewTriggerSORInfoUpdateAPIController(TriggerSORInfoUpdateAPIService)

	UEContextInAMFDataRetrievalAPIService := openapi.NewUEContextInAMFDataRetrievalAPIService()
	UEContextInAMFDataRetrievalAPIController := openapi.NewUEContextInAMFDataRetrievalAPIController(UEContextInAMFDataRetrievalAPIService)

	UEContextInSMFDataRetrievalAPIService := openapi.NewUEContextInSMFDataRetrievalAPIService()
	UEContextInSMFDataRetrievalAPIController := openapi.NewUEContextInSMFDataRetrievalAPIController(UEContextInSMFDataRetrievalAPIService)

	UEContextInSMSFDataRetrievalAPIService := openapi.NewUEContextInSMSFDataRetrievalAPIService()
	UEContextInSMSFDataRetrievalAPIController := openapi.NewUEContextInSMSFDataRetrievalAPIController(UEContextInSMSFDataRetrievalAPIService)

	UserConsentSubscriptionDataRetrievalAPIService := openapi.NewUserConsentSubscriptionDataRetrievalAPIService()
	UserConsentSubscriptionDataRetrievalAPIController := openapi.NewUserConsentSubscriptionDataRetrievalAPIController(UserConsentSubscriptionDataRetrievalAPIService)

	V2XSubscriptionDataRetrievalAPIService := openapi.NewV2XSubscriptionDataRetrievalAPIService()
	V2XSubscriptionDataRetrievalAPIController := openapi.NewV2XSubscriptionDataRetrievalAPIController(V2XSubscriptionDataRetrievalAPIService)

	router := openapi.NewRouter(AccessAndMobilitySubscriptionDataRetrievalAPIController, Class5MBSSubscriptionDataRetrievalAPIController, EnhancedCoverageRestrictionDataRetrievalAPIController, GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController, GroupIdentifiersAPIController, LCSBroadcastAssistanceDataTypesRetrievalAPIController, LCSMobileOriginatedDataRetrievalAPIController, LCSPrivacyDataRetrievalAPIController, MultipleIdentifiersAPIController, ProseSubscriptionDataRetrievalAPIController, ProvidingAcknowledgementOfCAGUpdateAPIController, ProvidingAcknowledgementOfSNSSAIsUpdateAPIController, ProvidingAcknowledgementOfSteeringOfRoamingAPIController, ProvidingAcknowledgementOfUEParametersUpdateAPIController, RetrievalOfMultipleDataSetsAPIController, RetrievalOfSharedDataAPIController, RetrievalOfTheIndividualSharedDataAPIController, SMFSelectionSubscriptionDataRetrievalAPIController, SMSManagementSubscriptionDataRetrievalAPIController, SMSSubscriptionDataRetrievalAPIController, SessionManagementSubscriptionDataRetrievalAPIController, SliceSelectionSubscriptionDataRetrievalAPIController, SubscriptionCreationAPIController, SubscriptionCreationForSharedDataAPIController, SubscriptionDeletionAPIController, SubscriptionDeletionForSharedDataAPIController, SubscriptionModificationAPIController, TraceConfigurationDataRetrievalAPIController, TriggerSORInfoUpdateAPIController, UEContextInAMFDataRetrievalAPIController, UEContextInSMFDataRetrievalAPIController, UEContextInSMSFDataRetrievalAPIController, UserConsentSubscriptionDataRetrievalAPIController, V2XSubscriptionDataRetrievalAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}