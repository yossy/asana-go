# Link asana action

Link Github and asana in Github actions

## Feature

- GitHubのPullRequestのURLを該当のAsanaのタスクの説明欄に書き込む
- GitHubActionsのjobで実行されることを想定しています。
- 連携したいRepogitoryのGithubActionsのWorkflowsから呼び出します。

## Usage

### 導入手順

- AsanaのDeveloperConsoleからPersonalAccessTokenを発行する
  - [AsanaのPersonalAccessTokenの発行について](https://asana.com/ja/guide/help/api/api#gl-access-tokens)
- 連携したいGitHubのRepogitoryのSettingsからSecretsに上記のAccessTokenを「`ASANA_TOKEN`」として登録する
- GitHubのSettings＞Developer settingsからPersonalAccessTokenを発行する
- 連携したいGitHubのrepogitoryのSettingsからSecretsにGitHubのAccessTokenを「`LINK_ASANA_ACTION_PAT`」として登録する
- `.github/workflows/update-asana-task.yml` を作成する
- `.github/pull_request_template.md` を作成し、以下を追加する

```md
<!-- 以下の()内にサブタスクのURLを貼ると自動でasana側タスクにPRのURLが追記されます -->

[Link Asana Task]()
```

- PRを作成する際に上記の `[Link Asana Task]()` の `（）` 内に紐付けたいAsanaのURLを入れる。

## Example usage

`.github/workflows/update-asana-task.yml`

```yml
name: Update Asana Task

on:
  pull_request:
    types: [ opened, edited ]

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
        repository: yossy/link-asana-action
        token: ${{ secrets.LINK_ASANA_ACTION_PAT }} # your github personal access token
    - name: Build package
      run: make build
    - name: Update asana task
      env:
        ASANA_TOKEN: ${{ secrets.ASANA_TOKEN }}
        PR: ${{ github.event.pull_request.html_url }}
        BODY: ${{ github.event.pull_request.body }}
      run: bin/updatetask -pat "$ASANA_TOKEN" -pr "$PR" -body "$BODY"
```
