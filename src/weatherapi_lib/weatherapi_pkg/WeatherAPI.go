/*
 * weatherapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package WeatherAPIClient

import(
	"weatherapi_lib/configuration_pkg"
	"weatherapi_lib/apis_pkg"
)


/*
 * Interface for the WEATHERAPI_IMPL
 */
type WEATHERAPI interface {
    APIs()                  apis_pkg.APIS
    Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the WEATHERAPI interface returning WEATHERAPI_IMPL
 */
func NewWEATHERAPI() WEATHERAPI {
    weatherAPIClient := new(WEATHERAPI_IMPL)
    weatherAPIClient.config = configuration_pkg.NewCONFIGURATION()

    return weatherAPIClient
}
