package operators

import(
	_"strings"
	"github.com/jptosso/coraza/pkg/models"
)

type BeginsWith struct{
	data string
	dlen int
}

func (o *BeginsWith) Init(data string){
	o.data = data
	o.dlen = len(data)
}

func (o *BeginsWith) Evaluate(tx *models.Transaction, value string) bool{
	if len(value) < o.dlen{
		return false
	}
	return o.data == value[0:o.dlen]
	//return strings.HasPrefix(value, o.data)
}