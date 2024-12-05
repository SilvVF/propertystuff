package main

type AtlasDeedResponse struct {
	SearchType   string `json:"search_type"`
	SearchParams struct {
	} `json:"search_params"`
	Query      string `json:"query"`
	Normalized string `json:"normalized"`
	Crs        struct {
		Type       string `json:"type"`
		Properties struct {
			Type string `json:"type"`
			Href string `json:"href"`
		} `json:"properties"`
	} `json:"crs"`
	Page      int    `json:"page"`
	PageCount int    `json:"page_count"`
	PageSize  int    `json:"page_size"`
	TotalSize int    `json:"total_size"`
	Type      string `json:"type"`
	Features  []struct {
		Type           string `json:"type"`
		AisFeatureType string `json:"ais_feature_type"`
		MatchType      string `json:"match_type"`
		Properties     struct {
			StreetAddress                 string   `json:"street_address"`
			AddressLow                    int      `json:"address_low"`
			AddressLowSuffix              string   `json:"address_low_suffix"`
			AddressLowFrac                string   `json:"address_low_frac"`
			AddressHigh                   any      `json:"address_high"`
			StreetPredir                  string   `json:"street_predir"`
			StreetName                    string   `json:"street_name"`
			StreetSuffix                  string   `json:"street_suffix"`
			StreetPostdir                 string   `json:"street_postdir"`
			UnitType                      string   `json:"unit_type"`
			UnitNum                       string   `json:"unit_num"`
			StreetFull                    string   `json:"street_full"`
			StreetCode                    int      `json:"street_code"`
			SegID                         int      `json:"seg_id"`
			ZipCode                       string   `json:"zip_code"`
			Zip4                          string   `json:"zip_4"`
			UspsBldgfirm                  string   `json:"usps_bldgfirm"`
			UspsType                      string   `json:"usps_type"`
			ElectionBlockID               string   `json:"election_block_id"`
			ElectionPrecinct              string   `json:"election_precinct"`
			PwdParcelID                   string   `json:"pwd_parcel_id"`
			DorParcelID                   string   `json:"dor_parcel_id"`
			LiAddressKey                  string   `json:"li_address_key"`
			EclipseLocationID             string   `json:"eclipse_location_id"`
			Bin                           string   `json:"bin"`
			LiParcelID                    string   `json:"li_parcel_id"`
			ZoningDocumentIds             []any    `json:"zoning_document_ids"`
			PwdAccountNums                []string `json:"pwd_account_nums"`
			OpaAccountNum                 string   `json:"opa_account_num"`
			OpaOwners                     []string `json:"opa_owners"`
			OpaAddress                    string   `json:"opa_address"`
			CenterCityDistrict            string   `json:"center_city_district"`
			CuaZone                       string   `json:"cua_zone"`
			LiDistrict                    string   `json:"li_district"`
			PhillyRisingArea              string   `json:"philly_rising_area"`
			CensusTract2010               string   `json:"census_tract_2010"`
			CensusBlockGroup2010          string   `json:"census_block_group_2010"`
			CensusBlock2010               string   `json:"census_block_2010"`
			CensusTract2020               string   `json:"census_tract_2020"`
			CensusBlockGroup2020          string   `json:"census_block_group_2020"`
			CensusBlock2020               string   `json:"census_block_2020"`
			CouncilDistrict2016           string   `json:"council_district_2016"`
			CouncilDistrict2024           string   `json:"council_district_2024"`
			PoliticalWard                 string   `json:"political_ward"`
			PoliticalDivision             string   `json:"political_division"`
			StateHouseRep2012             string   `json:"state_house_rep_2012"`
			StateHouseRep2022             string   `json:"state_house_rep_2022"`
			StateSenate2012               string   `json:"state_senate_2012"`
			StateSenate2022               string   `json:"state_senate_2022"`
			UsCongressional2012           string   `json:"us_congressional_2012"`
			UsCongressional2018           string   `json:"us_congressional_2018"`
			UsCongressional2022           string   `json:"us_congressional_2022"`
			PlanningDistrict              string   `json:"planning_district"`
			ElementarySchool              string   `json:"elementary_school"`
			MiddleSchool                  string   `json:"middle_school"`
			HighSchool                    string   `json:"high_school"`
			Zoning                        string   `json:"zoning"`
			ZoningRco                     string   `json:"zoning_rco"`
			CommercialCorridor            string   `json:"commercial_corridor"`
			HistoricDistrict              string   `json:"historic_district"`
			HistoricSite                  string   `json:"historic_site"`
			PoliceDivision                string   `json:"police_division"`
			PoliceDistrict                string   `json:"police_district"`
			PoliceServiceArea             string   `json:"police_service_area"`
			RubbishRecycleDay             string   `json:"rubbish_recycle_day"`
			SecondaryRubbishDay           string   `json:"secondary_rubbish_day"`
			RecyclingDiversionRate        float64  `json:"recycling_diversion_rate"`
			LeafCollectionArea            string   `json:"leaf_collection_area"`
			SanitationArea                string   `json:"sanitation_area"`
			SanitationDistrict            string   `json:"sanitation_district"`
			SanitationConvenienceCenter   string   `json:"sanitation_convenience_center"`
			CleanPhillyBlockCaptain       string   `json:"clean_philly_block_captain"`
			HistoricStreet                string   `json:"historic_street"`
			HighwayDistrict               string   `json:"highway_district"`
			HighwaySection                string   `json:"highway_section"`
			HighwaySubsection             string   `json:"highway_subsection"`
			TrafficDistrict               string   `json:"traffic_district"`
			TrafficPmDistrict             string   `json:"traffic_pm_district"`
			StreetLightRoute              string   `json:"street_light_route"`
			LaneClosure                   string   `json:"lane_closure"`
			PwdMaintDistrict              string   `json:"pwd_maint_district"`
			PwdPressureDistrict           string   `json:"pwd_pressure_district"`
			PwdTreatmentPlant             string   `json:"pwd_treatment_plant"`
			PwdWaterPlate                 string   `json:"pwd_water_plate"`
			PwdCenterCityDistrict         string   `json:"pwd_center_city_district"`
			MajorPhilaWatershed           string   `json:"major_phila_watershed"`
			NeighborhoodAdvisoryCommittee string   `json:"neighborhood_advisory_committee"`
			PprFriends                    string   `json:"ppr_friends"`
			EngineLocal                   string   `json:"engine_local"`
			LadderLocal                   string   `json:"ladder_local"`
			TobaccoRetailerPermitCapped   string   `json:"tobacco_retailer_permit_capped"`
			TobaccoFreeSchoolZones        string   `json:"tobacco_free_school_zones"`
		} `json:"properties"`
		Geometry struct {
			GeocodeType string    `json:"geocode_type"`
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

type AtlasDocuments struct {
	ObjectIDFieldName string `json:"objectIdFieldName"`
	UniqueIDField     struct {
		Name               string `json:"name"`
		IsSystemMaintained bool   `json:"isSystemMaintained"`
	} `json:"uniqueIdField"`
	GlobalIDFieldName string `json:"globalIdFieldName"`
	GeometryType      string `json:"geometryType"`
	SpatialReference  struct {
		Wkid       int `json:"wkid"`
		LatestWkid int `json:"latestWkid"`
	} `json:"spatialReference"`
	Fields []struct {
		Name         string `json:"name"`
		Type         string `json:"type"`
		Alias        string `json:"alias"`
		SQLType      string `json:"sqlType"`
		Domain       any    `json:"domain"`
		DefaultValue any    `json:"defaultValue"`
		Length       int    `json:"length,omitempty"`
	} `json:"fields"`
	Features []struct {
		Attributes struct {
			DocumentID   int    `json:"DOCUMENT_ID"`
			DisplayDate  int64  `json:"DISPLAY_DATE"`
			DocumentType string `json:"DOCUMENT_TYPE"`
			Grantors     string `json:"GRANTORS"`
			Grantees     string `json:"GRANTEES"`
			UnitNum      any    `json:"UNIT_NUM"`
		} `json:"attributes"`
	} `json:"features"`
}
