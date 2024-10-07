package song

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/task/internal/models"
)

// @Tags		Music
// @Param		songId	path	int		true	"songId"
// @Param		couplet	query	string	false	"Couplet"
// @Accept		json
// @Produce	json
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/song/{songId} [get]
func (sc *SongController) GetByID(c *gin.Context) {
	songId, _ := strconv.Atoi(c.Param("songId"))
	logFields := log.Fields{
		"requestType": "GET",
		"endpoint":    fmt.Sprintf("/song/%v", songId),
	}

	couplet, _ := strconv.Atoi(c.Query("couplet"))

	log.WithFields(logFields).Infof("request from user %+v:", couplet)

	song, err := sc.SongRepository.GetByID(c, uint(songId))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: models.ErrorDetail{
				Code:    "400",
				Message: "Get song error",
			},
		})
		log.WithFields(logFields).Errorf("Get song error:", err.Error())
		return
	}
	if couplet != 0 {
		song.Lyric = strings.Split(song.Lyric, "\n\n")[couplet-1]
	}
	log.WithFields(logFields).Infof("response from server: %+v", song)
	c.JSON(http.StatusOK, models.SuccessResponse{Result: song})
}
