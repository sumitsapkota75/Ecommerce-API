## E-commerce-API :books:

## How to get started ? :tada:
:arrow_right: create below two files in `root` directory of project :heavy_check_mark:

- `.env` <br />
- `serviceAccountKey.json` 

#### Reference documents :book:
1. https://drive.google.com/drive/folders/1L1z_SEj1CyvHkhGjranRYglcrq_EESaB?usp=sharing  

## commands :wrench:

| command             | desc                          |
| ------------------- | ------------------------------|
| `docker-compose up` | runs app                    |

#### Migration Commands :truck:

| Command             | desc                                           |
| ------------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |

