package users

//go:generate buf generate

//go:generate mockery --quiet --dir ./userspb -r --all --inpackage --case underscore
//go:generate mockery --quiet --dir ./internal -r --all --inpackage --case underscore
