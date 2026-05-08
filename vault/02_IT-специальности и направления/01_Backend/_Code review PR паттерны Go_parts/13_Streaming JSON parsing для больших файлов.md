---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]]"
type: concept
description: "data, _ := os.ReadFile(\"big.json\") // ← 5GB в память!"
---

## Streaming JSON parsing для больших файлов

### ❌ Антипаттерн

```go
data, _ := os.ReadFile("big.json")  // ← 5GB в память!
var items []Item
json.Unmarshal(data, &items)
```

### ✅ Streaming через json.Decoder

```go
file, _ := os.Open("big.json")
defer file.Close()

decoder := json.NewDecoder(file)
decoder.Token()  // [

for decoder.More() {
    var item Item
    if err := decoder.Decode(&item); err != nil {
        return err
    }
    process(item)  // обрабатываем одну запись за раз
}
```

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

