# Transmission

Social deduction game based on pairwise interaction.

## Dependencies

* [Git](https://git-scm.com/)
* [Go](https://golang.org/)
* [Make](https://www.gnu.org/software/make/)
* [Watchexec](https://github.com/watchexec/watchexec)
* [Yarn](https://yarnpkg.com/)

## Development

Open two terminals.

    $ make backend-dev
    $ make frontend-dev

Game is running at <http://localhost:3455/>.

## Production and deployment

Verify outside Docker.

    $ make backend-build
    $ make frontend-build
    $ make backend-prod

Verify inside Docker.

    $ make image-build
    $ make image-run

Deploy.

    $ make deploy
