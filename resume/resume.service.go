package resume

func GetByUsername(username string) (MyResume, error) {
	resume, err := FindByUsername(username)
	if err != nil {
		return MyResume{}, err
	}

	return resume, nil
}
