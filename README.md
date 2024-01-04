# go-clean-architecture
This repo is a demonstration on how to implement a clean architecture in Golang

Although there are many ways and variations on how to implement clean architecture, for this repo my goals are:

1. Write clean architecture in a Golang way. Using a common folder structure for Go projects (/cmd, /internal, /pkg...)
2. Implement it without many deviations or personal preferences. By the book. So that it can serve as a reference for those who want to understand clean architecture in the purest form.
3. Not be a complete scaffolding for new projects. Therefore, avoid placing a real DB connection, Docker structure... to avoid polluting it with too many files and information.
4. Trying to explicit the SOLID principles more clearly through the code organization.


Read more about [go repo organization](https://github.com/golang-standards/project-layout/tree/master).

Read more about [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

For a detailed explanation of the architecture, how it relates to SOLID, and common alternatives, you can read [this post](https://uian-sol.ghost.io/how-to-implement-clean-architecture-with-golang/).