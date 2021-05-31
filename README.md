# Travel-Agency-API

create below two files in root directory of project
1 .env
2 .serviceAccountKey.json 

Files reference to: https://drive.google.com/drive/folders/1L1z_SEj1CyvHkhGjranRYglcrq_EESaB?usp=sharing

`docker-compose up` -> to run project

#### Migration Commands

| Command            | Desc                                           |
| -------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |

