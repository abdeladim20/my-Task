social-network/
├── backend/
│   ├── pkg/
│   │   ├── db/
│   │   │   ├── migrations/
│   │   │   │   └── sqlite/
│   │   │   │       ├── 000001_create_users_table.up.sql
│   │   │   │       ├── 000001_create_users_table.down.sql
│   │   │   │       ├── 000002_create_posts_table.up.sql
│   │   │   │       └── 000002_create_posts_table.down.sql
│   │   │   └── sqlite.go
│   │   └── models/
│   │       ├── user.go
│   │       ├── post.go
│   │       └── comment.go
│   ├── routes/
│   │   ├── postRoutes.go
│   │   └── commentRoutes.go
│   └── server.go
│
└── frontend/
    ├── public/
    │   └── index.html
    ├── src/
    │   ├── assets/
    │   ├── components/
    │   │   ├── PostForm.vue
    │   │   ├── PostList.vue
    │   │   └── CommentForm.vue
    │   └── main.js
    └── package.json
