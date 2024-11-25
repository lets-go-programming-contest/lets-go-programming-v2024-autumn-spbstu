package contact

type ContactDB interface {
	AddContact(contact Contact) error
	GetContact(id int) error
}

type ContactRepository struct {
	ContactDB
}

func NewContactRepo(repo ContactDB) *ContactRepository {
	return &ContactRepository{repo}
}

func (r *ContactRepository) GetAll() ([]Contact, error) {

}

func (r *ContactRepository) GetByID(id int) (Contact, error) {

}

func (r *ContactRepository) Add(contact Contact) error {

}

func (r *ContactRepository) Update(id int, contact Contact) error {

}

func (r *ContactRepository) Delete(id int) error {

}
