package general

/*
	created by Ali Sadikin
	this file will provide all general function that all service (controller,model,repo) can use it

	## Function List
	DoResponse(params)
	WriteResponse(params)
	GetMetaData(params)
	SetQueryParams(param)
*/
import (
	"log"
	"os"
	"strconv"
)

/*
	-- this variable display all custom response code and its message
*/
var CustomResponseCode = map[string]string{
	// middle prefix 1 used for email accounts
	"010": "Email Not Found",
	"011": "Email Already Exists",
	"012": "User with this Email is Not Active",
	"013": "User with this Email does not exist",

	// middle prefix 2 for password
	"020": "Login Success",
	"021": "Password Doesn't Match",
	"022": "Error Creating Password",

	// first prefix 2 for record/database
	"030": "Record Succesfully Inserted",
	"031": "Record Succesfully Updated",
	"032": "Record Succesfully Deleted / Softdelete",
	"033": "Data Not Found",
	"034": "Data Found",
}

type Response struct {
	CC       string      `json:"ResponseCode"`
	Message  interface{} `json:"Message"`
	Data     interface{} `json:"Data"`
	MetaData Meta        `json:"Meta"`
}

type QueryParams struct {
	Page   uint8 `json:"Page"`
	Limit  uint8 `json:"Limit"`
	Status uint8 `json:"Status"`
}

type Meta struct {
	Page Pagination `json:"Pagination"`
}

type Pagination struct {
	Count uint8  `json:"Count"`
	URL   string `json:"URL"`
	// FirstPage   string `json:"FirstPage"`
	// CurrentPage string `json:"CurrentPage"`
	// NextPage    string `json:"NextPage"`
	// LastPage    string `json:"LastPage"`
}

/*
	-- func DoResponse(params) will display json data when invoked (generally) by controller
*/
func DoResponse(customeResposeCode string, message interface{}, data interface{}, meta interface{}) Response {
	// assertion
	assertedMeta := meta.(Meta)
	res := Response{CC: customeResposeCode, Message: message, Data: data, MetaData: assertedMeta}
	return res
}

/*
	-- func WriteErrorLog(param) will write error log and save it in to error.log file
*/
func WriteErrorLog(errorText string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(errorText + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

/*
	-- func GetMetaData(params) will display meta data / pagination on json response
*/
func GetMetaData(PageCount uint8, PageURL string) interface{} {
	if PageCount < 1 {
		PageURL = "Not Available"
	}
	meta := Meta{
		Page: Pagination{
			Count: PageCount,
			URL:   PageURL,
			// FirstPage:   FirstPage,
			// CurrentPage: CurrentPage,
			// NextPage:    NextPage,
			// LastPage:    LastPage,
		},
	}

	return meta
}

/*
	-- func SetQueryParams(param) will set query string params
*/
func SetQueryParams(allKeys map[string]interface{}) interface{} {
	keys := ""
	mark := "?"
	index := 0
	for key, val := range allKeys {
		if index > 0 {
			mark = "&"
		}

		index++
		keys = keys + mark + key + "=" + val.(string)
	}

	return keys
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
