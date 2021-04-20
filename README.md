# Asana link action

Link Github and asana in Github actions

## Example usage

```yml
name: Update Asana Task

  on:
    pull_request:
      types: [ edited ]

  jobs:
    asana:
      runs-on: ubuntu-latest
      steps:
      - name: Install Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.16
      - name: Checkout code
        uses: actions/checkout@v2
        with:
          repository: yossy/asana-go
      - name: Build package
        run: make build
      - name: Update asana task
        env:
          ASANA_TOKEN: ${{ secrets.ASANA_TOKEN }}
          PR: ${{ github.event.pull_request.html_url }}
          BODY: ${{ github.event.pull_request.body }}
        run: bin/updatetask -pat "$ASANA_TOKEN" -pr "$PR" -body "$BODY"
```
