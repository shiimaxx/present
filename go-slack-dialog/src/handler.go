var message slack.InteractionCallback
if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
    log.Print("json unmarshal message failed: ", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
}
//  :
if message.Type == "interactive_message" {
    // Order Coffeeボタンの入力を受け付けてフォームを表示する
} else if message.Type == "dialog_submission" {
    // フォームの入力を受け付けて何かする
}
