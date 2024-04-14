package repository

import "github.com/go-final-homework/models"

func (r *Repository) AddTransaction(transaction models.Transaction) (int, error) {
	_, err := r.db.Exec(
		"INSERT INTO transactions (amount, currency, type, category, description, date, user_id) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id",
		transaction.Amount,
		transaction.Currency,
		transaction.Type,
		transaction.Category,
		transaction.Description,
		transaction.Date,
		transaction.UserId)
	if err != nil {
		return 0, err
	}

	return transaction.Id, nil
}

func (r *Repository) GetAllTransactions() ([]models.Transaction, error) {
	rows, err := r.db.Query("SELECT * FROM transactions")
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	for rows.Next() {
		transaction := new(models.Transaction)
		err = rows.Scan(
			&transaction.Id,
			&transaction.Amount,
			&transaction.Currency,
			&transaction.Type,
			&transaction.Category,
			&transaction.Description,
			&transaction.Date,
			&transaction.UserId)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, *transaction)
	}
	return transactions, nil
}

func (r *Repository) GetTransactionById(id int) (models.Transaction, error) {
	rows, err := r.db.Query("SELECT * FROM transactions WHERE id = ?", id)
	if err != nil {
		return models.Transaction{}, err
	}
	var transaction models.Transaction
	for rows.Next() {
		err = rows.Scan(
			&transaction.Id,
			&transaction.Amount,
			&transaction.Currency,
			&transaction.Type,
			&transaction.Category,
			&transaction.Description,
			&transaction.Date,
			&transaction.UserId)
		if err != nil {
			return models.Transaction{}, err
		}
	}
	return transaction, nil
}

func (r *Repository) UpdateTransaction(id int, transaction models.Transaction) error {
	_, err := r.db.Exec(
		"UPDATE transactions SET amount = ?, currency = ?, type = ?, category = ?, description = ?, date = ?, user_id = ? WHERE id = ?",
		transaction.Amount,
		transaction.Currency,
		transaction.Type,
		transaction.Category,
		transaction.Description,
		transaction.Date,
		transaction.UserId,
		id)
	if err != nil {
		return err
	}	
	return nil
}

func (r *Repository) DeleteTransaction(id int) error {
	_, err := r.db.Exec("DELETE FROM transactions WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}