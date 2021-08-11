package main
import (
	
	"fmt"

	sq "github.com/Masterminds/squirrel"
	
	

)
var partyColumns = []string{`id`, `name`, `legacy_email`, `created_date`, `legacy_id`, `can_bid`, `death_date`,
	`preferred`, `preferred_tier`, `client_type`, `legacy_client_level`, `preferred_language_code`, `preferred_language_salutation`, `salutation`,
	`marketing_segment`, `client_status`, `country_code`, `updated_date`, `deceased`, `client_team_updated_date`, `associations_updated_date`,
	`phone_updated_date`, `address_updated_date`, `collection_value_low`, `collection_value_high`, `collection_value_date`, `collection_value_currency_code`,
	`likely_disposition`, `created_by`, `updated_by`, `associated_to_web`, `client_level`, `overview`, `client_sentiment`}
func main () {

	dateOfBirth := "200"
cc := make([]string, len(partyColumns))
for i, v := range partyColumns {
	cc[i] = fmt.Sprintf("party.%s", v)
}
b := sq.Select(cc...).From("party")
b = b.LeftJoin("individual ON party.id = individual.party_id").Limit(uint64(10)).OrderBy("party.id DESC")		
b = b.Where("date_of_birth=?", dateOfBirth)

b = b.Where("party.id > ?", "test")

query, args, err := b.ToSql()
fmt.Println(query)
fmt.Println("args")
fmt.Println(args)

if err != nil {
	 fmt.Errorf("failed to get sql: %w", err)
}
}

