package {{ required "Backage" .Backage }}

import (
    "time"

    "github.com/uptrace/bun"
    "github.com/google/uuid"
)

type {{ upperFirst .Model }} struct {
  ID uuid.UUID `bun:",pk"`

  CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
  UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
  DeletedAt time.Time `bun:",soft_delete"`
}

var _ bun.BeforeInsertHook = (*{{ upperFirst .Model }})(nil)

func ({{ firstLetter (lowerFirst .Model) }} *{{ upperFirst .Model }}) BeforeInsert() error {
    if u.ID == uuid.Nil {
        u.ID = uuid.New()
    }

    return nil
}
