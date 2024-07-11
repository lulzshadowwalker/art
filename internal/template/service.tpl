package {{ required "package name" .Backage }}

type (
    {{ required "name" .name | lowerFirst }}Service struct {
        repository Repository
    } 

    {{ upperFirst .name }}Repository interface {
        // 
    }
)

func New{{ upperFirst .name }}Service(r Repository) {{ lowerFirst .name }}Service {
    return {{ lowerFirst .name }}Service{
        repository: r, 
    }
}