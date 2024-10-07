package song

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/task/internal/models"
)

//	@Tags		Music
//	@Accept		json
//	@Produce	json
//	@Param		songId	path		int	true	"songId"
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/song/{songId} [delete]
func (sc *SongController) DeleteSong(c *gin.Context) {
	songId, _ := strconv.Atoi(c.Param("songId"))
	logFields := log.Fields{
		"requestType": "DELETE",
		"endpoint":    fmt.Sprintf("/song/%v", songId),
	}

	err := sc.SongRepository.DeleteSong(c, uint(songId))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: models.ErrorDetail{
				Code:    "400",
				Message: "Delete song error",
			},
		})
		log.WithFields(logFields).Errorf("Delete song error:", err.Error())
		return
	}
	log.WithFields(logFields).Infof("response from server: %+v", true)
	c.JSON(http.StatusOK, models.SuccessResponse{Result: true})
}
