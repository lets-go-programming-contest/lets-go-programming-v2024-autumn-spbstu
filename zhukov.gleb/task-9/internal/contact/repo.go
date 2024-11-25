package contact

type ContactRepository interface {
	GetContact() (id string, password []byte, err error)
	AddContact() error
}

type ContactMemoryRepository struct {
	Repo ContactRepository
}

func NewContactRepo(repo ContactRepository) *ContactMemoryRepository {
	return &ContactMemoryRepository{
		Repo: repo,
	}
}

func (r *ContactMemoryRepository) GetAll() ([]Contact, error) {

}

func (r *ContactMemoryRepository) GetByID(id uint32) (Contact, error) {

}

func (r *ContactMemoryRepository) Add(contact Contact) error {

}

func (r *ContactMemoryRepository) Update(id uint32, contact Contact) error {

}

func (r *ContactMemoryRepository) Delete(id uint32) error {

}
