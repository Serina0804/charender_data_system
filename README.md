# calendar_data_system

## 作るもの

スケジュール管理 + 振り返り記録アプリ

## 動機

1. スケジュール管理をしたい
    1. notion などは自分で sort しないといけない
    2. google calendar などではスクロールしないといけない + 日程を自分で探さないといけない
    3. パッと見て何の予定があるかが一目でわかるようなサービスが欲しい
2. 日々の記録を取るのが難しい
    1. notion などでは毎回自分で項目を追加しないといけない
3. スケジュール管理をしつつ，記録を取ることを，アプリを通して手軽に行えるようにしたい

   e.g, テニス: 普段ただ練習に参加するだけでは練習での気づきを記録するのが億劫になる．

   → このアプリを使うことでスケジュール管理のついでに，振り返りをすることができる


## 使うもの

- Go 言語

## 仕様

1. 予定入力画面（登録ボタンを押して登録する．）
2. 予定を入力したら以下のような画面に遷移する．
    1. 予定を表示する時は sort して，日付順に表示する．

| date | day of the week | event | start time | end time | memo | record |
| --- | --- | --- | --- | --- | --- | --- |
| 2024/4/29 | Mon | アプリ開発 | 10:30 | 12:30 | 構造体の相談 |  |
| 2024/5/2 | Thr | ミーティング | 13:00 | 14:30 | 研究頑張るぞ…! |  |
1. 予定を編集したい時は，はじめの予定入力画面に遷移する．
    1. 編集したいところだけ入力する．
    2. もし値が入力されればその値で元の値を上書きし，そうでなければ（＝値が NULL であれば）値は変更しない．

   ※ date - end_time: NULL は受け付けられない．何かしら値を入れないといけない．

   ※ memo について
    - 編集画面の時，登録したい文章をまた 1 から打ち直す感じになる..?

## 動作様子
- 初回動作時
[new](https://github.com/user-attachments/assets/354909f8-0ec5-405e-b702-11193595affd)

- スケジュール入力時：追加したいスケジュールを記入　→　Addにチェック　→　Executeボタンを押す
[before_adding_entry](https://github.com/user-attachments/assets/e14d23b0-e00b-4ff6-9f04-2cc5c21a4351)

- スケジュール入力後: スケジュールが表示される
[after_adding_entry](https://github.com/user-attachments/assets/59623cb5-325a-43f7-9656-3420a43e4beb)

- 複数のスケジュール入力後: スケジュールが時系列順に表示される
[after_adding_entry_sort](https://github.com/user-attachments/assets/cc2df34a-abae-4e15-bb33-0afac5fd64b6)

- スケジュールの検索時：検索したい日付を記入　→　Searchにチェック　→　Executeボタンを押す
[search_schedule](https://github.com/user-attachments/assets/1b6b7cd8-7537-4016-8747-7a176bf746a6)

- スケジュールの検索後：検索した日付のスケジュールが表示される
[after_search_schedule](https://github.com/user-attachments/assets/89640842-71a3-4cf7-b850-a018a7393b72)

<!-- - エラーメッセージ: 正しくない入力を受け付けた時
[error_message](https://github.com/Serina0804/charender_data_system/assets/126635893/359770a2-6e86-41f3-926e-371040c6cdc7) -->


## その他付け足したい機能（時間があれば）

- hash tag
- 検索機能
- hash tag 毎のページ