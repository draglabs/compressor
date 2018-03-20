package routes



import (
	"dsound/controllers"
	"dsound/types"
	"dsound/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

//JamRouter struct, is the jam router
var jamRouter = MainRouter.Group(APIV + "jam/")

const (
	recordingsR = "recording/:id"
	jamNewR     = "new"
	joinR       = "join"
	upload      = "upload"
	details     = "details/:id"
	updateJamR  = "update"
)

// addToMainROuter func, will add all the jam routes
// to he main router
func addToMainRouter() {
	jamRouter.GET(recordingsR, recordings)
	jamRouter.GET(details, jamDetails)
	jamRouter.POST(jamNewR, newJam)
	jamRouter.POST(joinR, join)
	jamRouter.POST(upload, uploadAudioFile)
	jamRouter.POST(updateJamR, updateJam)
}

// new func, will give us a new jam regarless of the user having an
// active jam, if the user has an active jam it will be replaced by this one
func newJam(c *gin.Context) {
	pm, err := utils.ParseJam(c)
	if err != nil {
		c.JSON(400, types.ResponseMessage{M: "One or more params are missing"})
		return
	}
	jam, err := controllers.Jam.Create(pm)
	if err == nil {
		c.JSON(200, jam)
		return
	}
	c.JSON(500, types.ResponseMessage{M: "Unable to create Jam"})
}

// upload func, takes care of the uplaoding, and currently uploads the file to
// s3 bucket.
func uploadAudioFile(c *gin.Context) {
	para, err := utils.ParseUpload(c)
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "Something went wrong"})
		return
	}
	err = controllers.Jam.Upload(para)
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "Something went wrong"})
		return
	}
	c.JSON(200, types.ResponseMessage{M: "uploaded succesfuly"})
}

// join func, join a user into a jam.
func join(c *gin.Context) {
	para, err := utils.ParseJoinJam(c)
	if err != nil {
		fmt.Println("error parsing join jam " + err.Error())
		c.JSON(500, types.ResponseMessage{M: "One or more params are missing"})
		return
	}

	if jam, err := controllers.Jam.Join(para); err == nil {
		c.JSON(200, jam)
		return
	}
}

func jamDetails(c *gin.Context) {
	jam, err := controllers.Jam.Details(c.Param("id"))
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "Something when wrong, Error: " + err.Error()})
		return
	}
	c.JSON(200, jam)

}

// recordings func, will fetch all the recordings for a given jam id.
func recordings(c *gin.Context) {
	id := c.Param("id")
	recordings, err := controllers.Recordings(id)
	if err != nil {
		c.JSON(400, types.ResponseMessage{M: "No recordings for this jam " + id})
		return
	}
	c.JSON(200, recordings)
}
func updateJam(c *gin.Context) {
	para, err := utils.ParseUpdate(c)
	jam, err := controllers.Jam.Update(para)
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "something went wrong" + err.Error()})
		return
	}
	if err == nil {
		c.JSON(200, jam)
	}
}


//import (
//	"net/http"
//	"fmt"
//)
//
//package routes
//
//import (
//"compressor/controllers"
//"compressor/models"
//"encoding/json"
//"fmt"
//"net/http"
//)
//
//// ArchiveRouter struct
//type ArchiveRouter struct {
//}
//type responseMessage struct {
//	Message string `json:"message"`
//}
//
//// NewArchiveRouter func
//// represents the single router we have
//func NewArchiveRouter() ArchiveRouter {
//	return ArchiveRouter{}
//}
//
//// i need to put this middleware before the interface method ServeHTTP
//func middleware(next http.HandlerFunc) http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println(r.URL)
//		w.Header().Set("Content-Type", "application/json")
//		next.ServeHTTP(w, r)
//	})
//}
//
//// ServeHTTP func, is the interface conformance
//func (ar ArchiveRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "POST":
//		w.Header().Set("Content-Type", "application/json")
//		handleArchive(w, r)
//
//	default:
//		w.WriteHeader(http.StatusNotImplemented)
//		w.Write([]byte("method not implemented"))
//
//		return
//	}
//}
//
//// handleArchive func, is the handler for the
//// archive route, it fetches the jam by a given id
//// and sends it back to the user
//func handleArchive(w http.ResponseWriter, r *http.Request) {
//	jam, err := controllers.FetchJam(parseParams(r))
//	if err == nil {
//		er := json.NewEncoder(w).Encode(&jam)
//		if er != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			json.NewEncoder(w).Encode(responseMessage{"something when wrong: " + er.Error()})
//		}
//	}
//}
//
//// parseParams func, parses the params of the incoming
//// request and checks for simple validation
//func parseParams(r *http.Request) *models.ArchiveParam {
//	var para models.ArchiveParam
//	err := json.NewDecoder(r.Body).Decode(&para)
//	defer r.Body.Close()
//	if err == nil {
//		fmt.Println(para.JamID)
//		return &para
//	}
//
//	return nil
//}
//
//// Index func, main index handler
//// handles the main entry point
//// by default it sends a message in json
//// to tell the user she/he as riched the index
//func Index(w http.ResponseWriter, r *http.Request) {
//	json.NewEncoder(w).Encode(responseMessage{"Compressor homepage"})
//}
