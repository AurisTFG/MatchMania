package services

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/utils"

	"github.com/google/uuid"
)

type QueueService interface {
	JoinQueue(leagueId uuid.UUID, teamId uuid.UUID) error
	LeaveQueue(leagueId uuid.UUID, teamId uuid.UUID) error
	GetAllQueues() ([]*responses.QueueDto, error)
	GetQueueByID(uuid.UUID) (*responses.QueueDto, error)
	CreateQueue(*models.Queue) error
	SaveQueue(*models.Queue) error
	DeleteQueue(uuid.UUID) error
}

type queueService struct {
	queueRepository repositories.QueueRepository
	teamRepository  repositories.TeamRepository
}

func NewQueueService(
	queueRepository repositories.QueueRepository,
	teamRepository repositories.TeamRepository,
) QueueService {
	return &queueService{
		queueRepository: queueRepository,
		teamRepository:  teamRepository,
	}
}

func (s *queueService) JoinQueue(leagueId uuid.UUID, teamId uuid.UUID) error {
	team, err := s.teamRepository.FindById(teamId)
	if err != nil {
		return err
	}

	queue, err := s.queueRepository.GetByLeagueID(leagueId)
	if err != nil {
		queue = &models.Queue{
			LeagueId: leagueId,
			GameMode: "2v2",
			Teams:    []models.Team{},
		}
	}

	queue.Teams = append(queue.Teams, *team)

	if err = s.queueRepository.Save(queue); err != nil {
		return err
	}

	return nil
}

func (s *queueService) LeaveQueue(leagueId uuid.UUID, teamId uuid.UUID) error {
	team, err := s.teamRepository.FindById(teamId)
	if err != nil {
		return err
	}

	team.QueueId = nil

	if err = s.teamRepository.Save(team); err != nil {
		return err
	}

	return nil
}

func (s *queueService) GetAllQueues() ([]*responses.QueueDto, error) {
	queues, err := s.queueRepository.GetAll()
	if err != nil {
		return nil, err
	}

	queueDtos := utils.MustCopy[[]*responses.QueueDto](queues)

	return *queueDtos, nil
}

func (s *queueService) GetQueueByID(id uuid.UUID) (*responses.QueueDto, error) {
	queue, err := s.queueRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	queueDto := utils.MustCopy[responses.QueueDto](queue)

	return queueDto, nil
}

func (s *queueService) CreateQueue(queue *models.Queue) error {
	err := s.queueRepository.Create(queue)
	if err != nil {
		return err
	}

	return nil
}

func (s *queueService) SaveQueue(queue *models.Queue) error {
	err := s.queueRepository.Save(queue)
	if err != nil {
		return err
	}

	return nil
}

func (s *queueService) DeleteQueue(id uuid.UUID) error {
	queue, err := s.queueRepository.GetByID(id)
	if err != nil {
		return err
	}

	err = s.queueRepository.Delete(queue.Id)
	if err != nil {
		return err
	}

	return nil
}
