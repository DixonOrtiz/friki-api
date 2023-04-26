package addressrepo

func (r AddressRepository) HardDeleteByID(ID int) {
	r.DB.Exec("DELETE FROM addresses WHERE id = ?", ID)
}
