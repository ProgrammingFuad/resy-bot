package main

import (
	"encoding/json"
	"time"
)

type DiscordChannelIdManager struct {
	winner string
	loser  string
}

type BookingResponse struct {
	ResyToken     *string `json:"resy_token"`
	ReservationID int     `json:"reservation_id"`
}

type DetailsResponse struct {
	Cancellation struct {
		Credit struct {
			DateCutOff interface{} `json:"date_cut_off"`
		} `json:"credit"`
		Display struct {
			Policy []string `json:"policy"`
		} `json:"display"`
		Fee    interface{} `json:"fee"`
		Refund struct {
			DateCutOff time.Time `json:"date_cut_off"`
		} `json:"refund"`
	} `json:"cancellation"`
	Change struct {
		DateCutOff time.Time `json:"date_cut_off"`
	} `json:"change"`
	Config struct {
		AddOns               interface{} `json:"add_ons"`
		DoubleConfirmation   interface{} `json:"double_confirmation"`
		Features             interface{} `json:"features"`
		MenuItems            interface{} `json:"menu_items"`
		ServiceChargeOptions interface{} `json:"service_charge_options"`
	} `json:"config"`
	Locale struct {
		Currency string `json:"currency"`
	} `json:"locale"`
	Payment struct {
		Amounts struct {
			Items             []interface{} `json:"items"`
			ReservationCharge json.Number   `json:"reservation_charge"`
			Subtotal          json.Number   `json:"subtotal"`
			AddOns            json.Number   `json:"add_ons"`
			Quantity          json.Number   `json:"quantity"`
			ResyFee           json.Number   `json:"resy_fee"`
			ServiceFee        json.Number   `json:"service_fee"`
			ServiceCharge     struct {
				Amount int    `json:"amount"`
				Value  string `json:"value"`
			} `json:"service_charge"`
			Tax          float64 `json:"tax"`
			Total        float64 `json:"total"`
			Surcharge    float64 `json:"surcharge"`
			PricePerUnit float64 `json:"price_per_unit"`
		} `json:"amounts"`
		Comp   bool `json:"comp"`
		Config struct {
			Type string `json:"type"`
		} `json:"config"`
		Display struct {
			Balance struct {
				Value    string `json:"value"`
				Modifier string `json:"modifier"`
			} `json:"balance"`
			Buy struct {
				Action         string `json:"action"`
				AfterModifier  string `json:"after_modifier"`
				BeforeModifier string `json:"before_modifier"`
				Init           string `json:"init"`
				Value          string `json:"value"`
			} `json:"buy"`
			Description []interface{} `json:"description"`
		} `json:"display"`
		Options []struct {
			Amounts struct {
				PricePerUnit  float64 `json:"price_per_unit"`
				ResyFee       int     `json:"resy_fee"`
				ServiceFee    int     `json:"service_fee"`
				ServiceCharge struct {
					Amount int    `json:"amount"`
					Value  string `json:"value"`
				} `json:"service_charge"`
				Tax   float64 `json:"tax"`
				Total float64 `json:"total"`
			} `json:"amounts"`
		} `json:"options"`
	} `json:"payment"`
	User *struct {
		PaymentMethods []struct {
			ID           *int64 `json:"id"`
			IsDefault    bool   `json:"is_default"`
			ProviderID   int    `json:"provider_id"`
			ProviderName string `json:"provider_name"`
			Display      string `json:"display"`
			Type         string `json:"type"`
			ExpYear      int    `json:"exp_year"`
			ExpMonth     int    `json:"exp_month"`
			IssuingBank  string `json:"issuing_bank"`
		} `json:"payment_methods"`
	} `json:"user"`
	Venue struct {
		Config struct {
			AllowBypassPaymentMethod int `json:"allow_bypass_payment_method"`
			AllowMultipleResys       int `json:"allow_multiple_resys"`
			EnableInvite             int `json:"enable_invite"`
			EnableResypay            int `json:"enable_resypay"`
			HospitalityIncluded      int `json:"hospitality_included"`
		} `json:"config"`
		Content []struct {
			Attribution interface{} `json:"attribution"`
			Body        string      `json:"body"`
			Display     struct {
				Type string `json:"type"`
			} `json:"display"`
			Icon struct {
				URL string `json:"url"`
			} `json:"icon"`
			Locale struct {
				Language string `json:"language"`
			} `json:"locale"`
			Name  string      `json:"name"`
			Title interface{} `json:"title"`
		} `json:"content"`
		Location struct {
			Address1     string      `json:"address_1"`
			Address2     interface{} `json:"address_2"`
			Code         string      `json:"code"`
			Country      string      `json:"country"`
			CrossStreet1 string      `json:"cross_street_1"`
			CrossStreet2 string      `json:"cross_street_2"`
			ID           int         `json:"id"`
			Latitude     float64     `json:"latitude"`
			Locality     string      `json:"locality"`
			Longitude    float64     `json:"longitude"`
			Neighborhood string      `json:"neighborhood"`
			PostalCode   string      `json:"postal_code"`
			Region       string      `json:"region"`
		} `json:"location"`
		Rater []struct {
			Name  string  `json:"name"`
			Scale int     `json:"scale"`
			Score float64 `json:"score"`
			Total int     `json:"total"`
		} `json:"rater"`
		Source struct {
			Name           interface{} `json:"name"`
			Logo           interface{} `json:"logo"`
			TermsOfService interface{} `json:"terms_of_service"`
			PrivacyPolicy  interface{} `json:"privacy_policy"`
		} `json:"source"`
	} `json:"venue"`
	Viewers struct {
		Total int `json:"total"`
	} `json:"viewers"`
	BookToken struct {
		DateExpires time.Time `json:"date_expires"`
		Value       *string   `json:"value"`
	} `json:"book_token"`
}

type FindReservationResponse struct {
	Query struct {
		Day       string `json:"day"`
		PartySize int    `json:"party_size"`
	} `json:"query"`
	Bookmark interface{} `json:"bookmark"`
	Results  *struct {
		Venues []struct {
			Venue struct {
				ID struct {
					Resy int `json:"resy"`
				} `json:"id"`
				VenueGroup struct {
					ID     int           `json:"id"`
					Name   string        `json:"name"`
					Venues []interface{} `json:"venues"`
				} `json:"venue_group"`
				Name                         string  `json:"name"`
				Type                         string  `json:"type"`
				URLSlug                      string  `json:"url_slug"`
				PriceRange                   int     `json:"price_range"`
				AverageBillSize              float64 `json:"average_bill_size"`
				CurrencySymbol               string  `json:"currency_symbol"`
				HospitalityIncluded          int     `json:"hospitality_included"`
				ResySelect                   int     `json:"resy_select"`
				IsGdc                        int     `json:"is_gdc"`
				IsGlobalDiningAccess         bool    `json:"is_global_dining_access"`
				IsGlobalDiningAccessOnly     bool    `json:"is_global_dining_access_only"`
				RequiresReservationTransfers int     `json:"requires_reservation_transfers"`
				IsGns                        int     `json:"is_gns"`
				TransactionProcessor         string  `json:"transaction_processor"`
				HideAllergyQuestion          bool    `json:"hide_allergy_question"`
				HideOccasionQuestion         bool    `json:"hide_occasion_question"`
				HideSpecialRequestQuestion   bool    `json:"hide_special_request_question"`
				GdaConciergeBooking          bool    `json:"gda_concierge_booking"`
				TaxIncluded                  bool    `json:"tax_included"`
				Rating                       float64 `json:"rating"`
				TotalRatings                 int     `json:"total_ratings"`
				Inventory                    struct {
					Type struct {
						ID int `json:"id"`
					} `json:"type"`
				} `json:"inventory"`
				Reopen struct {
					Date string `json:"date"`
				} `json:"reopen"`
				Location struct {
					TimeZone     string `json:"time_zone"`
					Neighborhood string `json:"neighborhood"`
					Geo          struct {
						Lat float64 `json:"lat"`
						Lon float64 `json:"lon"`
					} `json:"geo"`
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"location"`
				TravelTime struct {
					Distance float64 `json:"distance"`
				} `json:"travel_time"`
				Source struct {
					Name           interface{} `json:"name"`
					Logo           interface{} `json:"logo"`
					TermsOfService interface{} `json:"terms_of_service"`
					PrivacyPolicy  interface{} `json:"privacy_policy"`
				} `json:"source"`
				ServiceTypes struct {
					Num2 struct {
					} `json:"2"`
				} `json:"service_types"`
				Top    bool `json:"top"`
				Ticket struct {
					Average    float64 `json:"average"`
					AverageStr string  `json:"average_str"`
				} `json:"ticket"`
				Currency struct {
					Symbol string `json:"symbol"`
					Code   string `json:"code"`
				} `json:"currency"`
				IsRga            bool   `json:"is_rga"`
				IsRgaOnly        bool   `json:"is_rga_only"`
				DefaultTemplate  string `json:"default_template"`
				ResponsiveImages struct {
					Originals struct {
						E77Fb9Cbe55B833D4A122F7293Cb96E3A468F386 struct {
							URL string `json:"url"`
						} `json:"e77fb9cbe55b833d4a122f7293cb96e3a468f386"`
						B4853650F239Cf66Ed30C10B7D22E38Bdcb5F9B5 struct {
							URL string `json:"url"`
						} `json:"b4853650f239cf66ed30c10b7d22e38bdcb5f9b5"`
						B40Ab1A45Ef6Ce74F92Ace31B4678637414Cd059 struct {
							URL string `json:"url"`
						} `json:"b40ab1a45ef6ce74f92ace31b4678637414cd059"`
						Five7Ead857E30B406E5A96C2Cc6E2F0C7E56832952 struct {
							URL string `json:"url"`
						} `json:"57ead857e30b406e5a96c2cc6e2f0c7e56832952"`
						TwoC34269F24Ccaf9Adac8C8011F40E88Db71Ec5Cc struct {
							URL string `json:"url"`
						} `json:"2c34269f24ccaf9adac8c8011f40e88db71ec5cc"`
					} `json:"originals"`
					Urls struct {
						E77Fb9Cbe55B833D4A122F7293Cb96E3A468F386 struct {
							One1 struct {
								Num200  string `json:"200"`
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"1:1"`
							Four3 struct {
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"4:3"`
							One69 struct {
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"16:9"`
						} `json:"e77fb9cbe55b833d4a122f7293cb96e3a468f386"`
						B4853650F239Cf66Ed30C10B7D22E38Bdcb5F9B5 struct {
							One1 struct {
								Num200 string `json:"200"`
								Num400 string `json:"400"`
								Num800 string `json:"800"`
							} `json:"1:1"`
							Four3 struct {
								Num400 string `json:"400"`
								Num800 string `json:"800"`
							} `json:"4:3"`
							One69 struct {
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"16:9"`
						} `json:"b4853650f239cf66ed30c10b7d22e38bdcb5f9b5"`
						B40Ab1A45Ef6Ce74F92Ace31B4678637414Cd059 struct {
							One1 struct {
								Num200 string `json:"200"`
								Num400 string `json:"400"`
								Num800 string `json:"800"`
							} `json:"1:1"`
							Four3 struct {
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"4:3"`
							One69 struct {
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"16:9"`
						} `json:"b40ab1a45ef6ce74f92ace31b4678637414cd059"`
						Five7Ead857E30B406E5A96C2Cc6E2F0C7E56832952 struct {
							One1 struct {
								Num200  string `json:"200"`
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"1:1"`
							Four3 struct {
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"4:3"`
							One69 struct {
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"16:9"`
						} `json:"57ead857e30b406e5a96c2cc6e2f0c7e56832952"`
						TwoC34269F24Ccaf9Adac8C8011F40E88Db71Ec5Cc struct {
							One1 struct {
								Num200 string `json:"200"`
								Num400 string `json:"400"`
								Num800 string `json:"800"`
							} `json:"1:1"`
							Four3 struct {
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"4:3"`
							One69 struct {
								Num400  string `json:"400"`
								Num800  string `json:"800"`
								Num1600 string `json:"1600"`
							} `json:"16:9"`
						} `json:"2c34269f24ccaf9adac8c8011f40e88db71ec5cc"`
					} `json:"urls"`
					UrlsByResolution struct {
						E77Fb9Cbe55B833D4A122F7293Cb96E3A468F386 struct {
							Num200 struct {
								One1 string `json:"1:1"`
							} `json:"200"`
							Num400 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"400"`
							Num800 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"800"`
							Num1600 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"1600"`
						} `json:"e77fb9cbe55b833d4a122f7293cb96e3a468f386"`
						B4853650F239Cf66Ed30C10B7D22E38Bdcb5F9B5 struct {
							Num200 struct {
								One1 string `json:"1:1"`
							} `json:"200"`
							Num400 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"400"`
							Num800 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"800"`
							Num1600 struct {
								One69 string `json:"16:9"`
							} `json:"1600"`
						} `json:"b4853650f239cf66ed30c10b7d22e38bdcb5f9b5"`
						B40Ab1A45Ef6Ce74F92Ace31B4678637414Cd059 struct {
							Num200 struct {
								One1 string `json:"1:1"`
							} `json:"200"`
							Num400 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"400"`
							Num800 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"800"`
							Num1600 struct {
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"1600"`
						} `json:"b40ab1a45ef6ce74f92ace31b4678637414cd059"`
						Five7Ead857E30B406E5A96C2Cc6E2F0C7E56832952 struct {
							Num200 struct {
								One1 string `json:"1:1"`
							} `json:"200"`
							Num400 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"400"`
							Num800 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"800"`
							Num1600 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"1600"`
						} `json:"57ead857e30b406e5a96c2cc6e2f0c7e56832952"`
						TwoC34269F24Ccaf9Adac8C8011F40E88Db71Ec5Cc struct {
							Num200 struct {
								One1 string `json:"1:1"`
							} `json:"200"`
							Num400 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"400"`
							Num800 struct {
								One1  string `json:"1:1"`
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"800"`
							Num1600 struct {
								Four3 string `json:"4:3"`
								One69 string `json:"16:9"`
							} `json:"1600"`
						} `json:"2c34269f24ccaf9adac8c8011f40e88db71ec5cc"`
					} `json:"urls_by_resolution"`
					FileNames    []string `json:"file_names"`
					AspectRatios struct {
						One1 struct {
							Num200  string `json:"200"`
							Num400  string `json:"400"`
							Num800  string `json:"800"`
							Num1600 string `json:"1600"`
						} `json:"1:1"`
						Four3 struct {
							Num400  string `json:"400"`
							Num800  string `json:"800"`
							Num1600 string `json:"1600"`
						} `json:"4:3"`
						One69 struct {
							Num400  string `json:"400"`
							Num800  string `json:"800"`
							Num1600 string `json:"1600"`
						} `json:"16:9"`
					} `json:"aspect_ratios"`
				} `json:"responsive_images"`
				NotifyOptions []struct {
					ServiceTypeID      int    `json:"service_type_id"`
					MinRequestDatetime string `json:"min_request_datetime"`
					MaxRequestDatetime string `json:"max_request_datetime"`
					StepMinutes        int    `json:"step_minutes"`
				} `json:"notify_options"`
				Favorite bool `json:"favorite"`
				Waitlist struct {
					Available int         `json:"available"`
					Label     string      `json:"label"`
					Current   interface{} `json:"current"`
				} `json:"waitlist"`
				SupportsPickups int           `json:"supports_pickups"`
				Collections     []interface{} `json:"collections"`
				Content         []struct {
					Attribution interface{} `json:"attribution"`
					Body        string      `json:"body"`
					Display     struct {
						Type string `json:"type"`
					} `json:"display"`
					Icon struct {
						URL string `json:"url"`
					} `json:"icon"`
					Locale struct {
						Language string `json:"language"`
					} `json:"locale"`
					Name  string      `json:"name"`
					Title interface{} `json:"title"`
				} `json:"content"`
				AllowBypassPaymentMethod int           `json:"allow_bypass_payment_method"`
				Events                   []interface{} `json:"events"`
			} `json:"venue"`
			Templates struct {
				Num1648000 struct {
					IsPaid                     bool          `json:"is_paid"`
					VenueShare                 interface{}   `json:"venue_share"`
					RestrictionID              interface{}   `json:"restriction_id"`
					PaymentStructure           interface{}   `json:"payment_structure"`
					CancellationFee            interface{}   `json:"cancellation_fee"`
					SecsCancelCutOff           interface{}   `json:"secs_cancel_cut_off"`
					TimeCancelCutOff           interface{}   `json:"time_cancel_cut_off"`
					SecsChangeCutOff           interface{}   `json:"secs_change_cut_off"`
					TimeChangeCutOff           interface{}   `json:"time_change_cut_off"`
					LargePartySizeCancel       interface{}   `json:"large_party_size_cancel"`
					LargePartyCancellationFee  interface{}   `json:"large_party_cancellation_fee"`
					LargePartySecsCancelCutOff interface{}   `json:"large_party_secs_cancel_cut_off"`
					LargePartyTimeCancelCutOff interface{}   `json:"large_party_time_cancel_cut_off"`
					LargePartySecsChangeCutOff interface{}   `json:"large_party_secs_change_cut_off"`
					LargePartyTimeChangeCutOff interface{}   `json:"large_party_time_change_cut_off"`
					DepositFee                 interface{}   `json:"deposit_fee"`
					ServiceCharge              interface{}   `json:"service_charge"`
					ServiceChargeOptions       []interface{} `json:"service_charge_options"`
					Images                     []string      `json:"images"`
					RawImageNames              []string      `json:"raw_image_names"`
					ImageDimensions            [][]int       `json:"image_dimensions"`
					IsDefault                  int           `json:"is_default"`
					IsEvent                    int           `json:"is_event"`
					IsPickup                   int           `json:"is_pickup"`
					PickupHighlight            int           `json:"pickup_highlight"`
					VenueID                    int           `json:"venue_id"`
					ReservationConfig          struct {
						Badge         interface{} `json:"badge"`
						Type          string      `json:"type"`
						SecsOffMarket int         `json:"secs_off_market"`
						TimeOffMarket interface{} `json:"time_off_market"`
					} `json:"reservation_config"`
					TurnTimes []struct {
						SecsAmount int `json:"secs_amount"`
						Size       struct {
							Max interface{} `json:"max"`
							Min int         `json:"min"`
						} `json:"size"`
					} `json:"turn_times"`
					DisplayConfig struct {
						Color struct {
							Background interface{} `json:"background"`
							Font       interface{} `json:"font"`
						} `json:"color"`
					} `json:"display_config"`
					Content struct {
						EnUs struct {
							About struct {
								Attribution interface{} `json:"attribution"`
								Body        string      `json:"body"`
								Title       interface{} `json:"title"`
							} `json:"about"`
						} `json:"en-us"`
					} `json:"content"`
					ID   int `json:"id"`
					Menu struct {
					} `json:"menu"`
					Name    string        `json:"name"`
					ItemIds []interface{} `json:"item_ids"`
					MenuIds []interface{} `json:"menu_ids"`
				} `json:"1648000"`
				Num1649264 struct {
					IsPaid                     bool          `json:"is_paid"`
					VenueShare                 interface{}   `json:"venue_share"`
					RestrictionID              interface{}   `json:"restriction_id"`
					PaymentStructure           interface{}   `json:"payment_structure"`
					CancellationFee            interface{}   `json:"cancellation_fee"`
					SecsCancelCutOff           interface{}   `json:"secs_cancel_cut_off"`
					TimeCancelCutOff           interface{}   `json:"time_cancel_cut_off"`
					SecsChangeCutOff           interface{}   `json:"secs_change_cut_off"`
					TimeChangeCutOff           interface{}   `json:"time_change_cut_off"`
					LargePartySizeCancel       interface{}   `json:"large_party_size_cancel"`
					LargePartyCancellationFee  interface{}   `json:"large_party_cancellation_fee"`
					LargePartySecsCancelCutOff interface{}   `json:"large_party_secs_cancel_cut_off"`
					LargePartyTimeCancelCutOff interface{}   `json:"large_party_time_cancel_cut_off"`
					LargePartySecsChangeCutOff interface{}   `json:"large_party_secs_change_cut_off"`
					LargePartyTimeChangeCutOff interface{}   `json:"large_party_time_change_cut_off"`
					DepositFee                 interface{}   `json:"deposit_fee"`
					ServiceCharge              interface{}   `json:"service_charge"`
					ServiceChargeOptions       []interface{} `json:"service_charge_options"`
					Images                     []string      `json:"images"`
					RawImageNames              []string      `json:"raw_image_names"`
					ImageDimensions            [][]int       `json:"image_dimensions"`
					IsDefault                  int           `json:"is_default"`
					IsEvent                    int           `json:"is_event"`
					IsPickup                   int           `json:"is_pickup"`
					PickupHighlight            int           `json:"pickup_highlight"`
					VenueID                    int           `json:"venue_id"`
					ReservationConfig          struct {
						Badge         interface{} `json:"badge"`
						Type          string      `json:"type"`
						SecsOffMarket interface{} `json:"secs_off_market"`
						TimeOffMarket interface{} `json:"time_off_market"`
					} `json:"reservation_config"`
					TurnTimes []struct {
						SecsAmount int `json:"secs_amount"`
						Size       struct {
							Max int `json:"max"`
							Min int `json:"min"`
						} `json:"size"`
					} `json:"turn_times"`
					DisplayConfig struct {
						Color struct {
							Background interface{} `json:"background"`
							Font       interface{} `json:"font"`
						} `json:"color"`
					} `json:"display_config"`
					Content struct {
						EnUs struct {
							About struct {
								Attribution interface{} `json:"attribution"`
								Body        string      `json:"body"`
								Title       interface{} `json:"title"`
							} `json:"about"`
						} `json:"en-us"`
					} `json:"content"`
					ID   int `json:"id"`
					Menu struct {
					} `json:"menu"`
					Name    string        `json:"name"`
					ItemIds []interface{} `json:"item_ids"`
					MenuIds []interface{} `json:"menu_ids"`
				} `json:"1649264"`
			} `json:"templates"`
			Slots []struct {
				Availability struct {
					ID int `json:"id"`
				} `json:"availability"`
				Config struct {
					ID    int    `json:"id"`
					Type  string `json:"type"`
					Token string `json:"token"`
				} `json:"config"`
				Date struct {
					End   string `json:"end"`
					Start string `json:"start"`
				} `json:"date"`
				Exclusive struct {
					ID int `json:"id"`
				} `json:"exclusive"`
				IsGlobalDiningAccess bool `json:"is_global_dining_access"`
				Floorplan            struct {
					ID int `json:"id"`
				} `json:"floorplan"`
				ID     interface{} `json:"id"`
				Market struct {
					Date struct {
						Off int `json:"off"`
						On  int `json:"on"`
					} `json:"date"`
				} `json:"market"`
				Meta struct {
					Size struct {
						Assumed int `json:"assumed"`
					} `json:"size"`
					Type struct {
						ID int `json:"id"`
					} `json:"type"`
				} `json:"meta"`
				Lock   interface{} `json:"lock"`
				Pacing struct {
					Beyond bool `json:"beyond"`
				} `json:"pacing"`
				Score struct {
					Total float64 `json:"total"`
				} `json:"score"`
				Shift struct {
					ID      int `json:"id"`
					Service struct {
						Type struct {
							ID int `json:"id"`
						} `json:"type"`
					} `json:"service"`
					Day string `json:"day"`
				} `json:"shift"`
				Size struct {
					Max int `json:"max"`
					Min int `json:"min"`
				} `json:"size"`
				Status struct {
					ID int `json:"id"`
				} `json:"status"`
				Table struct {
					ID []int `json:"id"`
				} `json:"table"`
				Template struct {
					ID int `json:"id"`
				} `json:"template"`
				Time struct {
					Turn struct {
						Actual    int `json:"actual"`
						Estimated int `json:"estimated"`
					} `json:"turn"`
				} `json:"time"`
				Quantity      int `json:"quantity"`
				DisplayConfig struct {
					Color struct {
						Background interface{} `json:"background"`
						Font       interface{} `json:"font"`
					} `json:"color"`
				} `json:"display_config"`
				ReservationConfig struct {
					Badge interface{} `json:"badge"`
				} `json:"reservation_config"`
				GdcPerk interface{} `json:"gdc_perk"`
				Payment struct {
					IsPaid               bool          `json:"is_paid"`
					CancellationFee      interface{}   `json:"cancellation_fee"`
					DepositFee           interface{}   `json:"deposit_fee"`
					ServiceCharge        interface{}   `json:"service_charge"`
					VenueShare           interface{}   `json:"venue_share"`
					PaymentStructure     interface{}   `json:"payment_structure"`
					SecsCancelCutOff     interface{}   `json:"secs_cancel_cut_off"`
					TimeCancelCutOff     interface{}   `json:"time_cancel_cut_off"`
					SecsChangeCutOff     interface{}   `json:"secs_change_cut_off"`
					TimeChangeCutOff     interface{}   `json:"time_change_cut_off"`
					ServiceChargeOptions []interface{} `json:"service_charge_options"`
				} `json:"payment"`
			} `json:"slots"`
			Notifies []interface{} `json:"notifies"`
			Pickups  struct {
				Slots        []interface{} `json:"slots"`
				ServiceTypes struct {
				} `json:"service_types"`
			} `json:"pickups"`
		} `json:"venues"`
		Meta struct {
			Offset int         `json:"offset"`
			Limit  interface{} `json:"limit"`
		} `json:"meta"`
	} `json:"results"`
}
