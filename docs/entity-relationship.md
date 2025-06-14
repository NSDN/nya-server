# Entity Relationship Diagram

使用 [Mermaid](https://mermaid.js.org/syntax/entityRelationshipDiagram.html) 绘制。

```Mermaid
erDiagram
  PLATES {
    VARCHAR id PK
    VARCHAR name
    TEXT description
    TEXT background
    VARCHAR page_type
    SMALLINT sort_order UK ">= 0"
  }
```
