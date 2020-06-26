# i18n-time

## Whats is this...
- Displays the current Local time & GMT(UTC) time. 
- World current time
- Specified Timezone current time

## How to use.
- Default(No ARG): "Local & GMT"
  ```Shell
    $ go run main.go
    Thu, 13 Feb 2020 23:47:58 JST   : Local
    Thu, 13 Feb 2020 14:47:58 GMT   : Etc/GMT
  ```

- i18n: World Timezone List
  ```Shell
    $ go run main.go i18n
    Thu, 13 Feb 2020 15:52:01 CET   : Europe/Andorra
    Thu, 13 Feb 2020 18:52:01 +04   : Asia/Dubai
    - snip -
    Thu, 13 Feb 2020 16:52:01 SAST  : Africa/Johannesburg
    Thu, 13 Feb 2020 16:52:01 CAT   : Africa/Lusaka
    Thu, 13 Feb 2020 16:52:01 CAT   : Africa/Harare
  ```

- "{TZ}": Specified Timezone 
  ```Shell
    $ go run main.go "{TZ}"
  ```
  - Ex: "America/Chicago"
    ```Shell
      $ go run main.go America/Chicago
      Thu, 13 Feb 2020 23:55:38 JST   : Local
      Thu, 13 Feb 2020 08:55:38 CST   : America/Chicago
    ```

- help: This message. :-)
  ```Shell
    $ go run main.go help
  ```
