package services

import (
	"chess-server/database"
	"chess-server/models"
	"log"
)

func CreateMatch(match models.Match) (*models.MatchResponse, error) {
	result := database.DB.Create(&match)
	if result.Error != nil {
		log.Println("Error al crear el partido:", result.Error)
		return nil, result.Error
	}
	log.Println("Partido creado:", match)

	response := models.MatchResponse{
		ID:        match.ID,
		CreatedAt: match.CreatedAt,
		WinnerId:  match.WinnerId,
		Player2id: match.Player2id,
		Player3id: match.Player3id,
		Player4id: match.Player4id,
		Time:      match.Time,
	}
	return &response, nil
}

func GetAllMatches() []models.MatchResponse {
	var matches []models.Match
	result := database.DB.Find(&matches)
	if result.Error != nil {
		log.Println("Error al obtener partidos:", result.Error)
	}
	log.Println("Partidos obtenidos:", matches)

	var matchesResponse []models.MatchResponse
	for _, match := range matches {
		matchResponse := models.MatchResponse{
			ID:        match.ID,
			CreatedAt: match.CreatedAt,
			WinnerId:  match.WinnerId,
			Player2id: match.Player2id,
			Player3id: match.Player3id,
			Player4id: match.Player4id,
			Time:      match.Time,
		}
		matchesResponse = append(matchesResponse, matchResponse)
	}
	return matchesResponse
}

func DeleteMatch(id uint) error {
	result := database.DB.Delete(&models.Match{}, id)
	if result.Error != nil {
		log.Println("Error al eliminar el partido:", result.Error)
		return result.Error
	}
	log.Println("Partido eliminado:", id)
	return nil
}

func GetMatchesByPlayerId(id uint) ([]models.MatchResponseUsername, error) {
    var matches []models.Match

    result := database.DB.Where("winner_id = ? OR player2id = ? OR player3id = ? OR player4id = ?", id, id, id, id).
        Find(&matches)

    if result.Error != nil {
        log.Println("Error al obtener las partidas:", result.Error)
        return nil, result.Error
    }

    var matchResponses []models.MatchResponseUsername
    for _, match := range matches {
        var winnerUsername, player2Username, player3Username, player4Username string

        database.DB.Model(&models.User{}).Where("id = ?", match.WinnerId).Pluck("username", &winnerUsername)
        database.DB.Model(&models.User{}).Where("id = ?", match.Player2id).Pluck("username", &player2Username)
        database.DB.Model(&models.User{}).Where("id = ?", match.Player3id).Pluck("username", &player3Username)
        database.DB.Model(&models.User{}).Where("id = ?", match.Player4id).Pluck("username", &player4Username)

        matchResponse := models.MatchResponseUsername{
            ID:              match.ID,
            CreatedAt:       match.CreatedAt,
            WinnerUsername:  winnerUsername,
            Player2Username: player2Username,
            Player3Username: player3Username,
            Player4Username: player4Username,
            Time:            match.Time,
        }
        matchResponses = append(matchResponses, matchResponse)
    }
    return matchResponses, nil
}
