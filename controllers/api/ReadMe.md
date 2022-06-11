
## API Reference

### Users

#### Get all users

```http
  GET /api/users
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get user

```http
  GET /api/user/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to fetch |

#### Create user

```http
  Post /api/user
```

| Parameter       | Type     | Description                                   |
| :--------       | :------- | :--------------------------------             |
| `username`      | `string` | **Required**. name of new user                |
| `password`      | `string` | **Required**. password  of new user           |
| `email`         | `string` | **Required**. email     of new user           |
| `firstname`     | `string` | **Required**. firstname of new user           |
| `lastname`      | `string` | **Required**. lastname  of new user           |
| `birthdate`     | `string` | **Required**. birthdate of new user           |
| `genre`         | `string` | **Required**. genre     of new user           |
| `bio`           | `string` | **Not Required**. bio     of new user         |
| `profilepicture`| `string` | **Not Required**. profilepicture of new user  |



 








