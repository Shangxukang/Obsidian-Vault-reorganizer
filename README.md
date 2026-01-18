# Obsidian-Vault-reorganizer

Obsidian以Markdown为语法记笔记，文件中的图片通常仅以`![[xxx.png]]` 或 `![[assets/xxx.png|300]]` 的形式存在，并不包含完整的文件路径。  

因此在**跨 Vault 迁移笔记**时，图片文件往往不会随笔记一起移动，导致迁移后的笔记中出现图片丢失、无法渲染的问题。

本工具用于**一键将旧 Vault 中对应的图片迁移到新 Vault 中**。

Obsidian uses Markdown syntax for note-taking. Images in notes are usually referenced as `![[xxx.png]]` or `![[assets/xxx.png|300]]`, without full file paths.

As a result, when **migrating notes across Vaults**, image files often do not move together with the notes, leading to missing images and rendering issues in the new Vault.

This tool is designed to **migrate the corresponding image files from the old Vault to the new Vault in one step**, ensuring that images are preserved and rendered correctly after migration.

---

## 用法 / Manual

先将markdown文件迁入新vault中。然后在终端中，以本项目目录为工作目录，运行：

``` bash
go run main.go 'new-vault-path' 'old-vault-path'
```
便完成迁移。

其中：

`new-vault-path`：新Vault的绝对路径

`old-vault-path`：旧Vault的绝对路径（图片可以在任何层级的子文件夹中）

注意path由单引号包裹。


## 示例 / Example
新 Vault:
 /Users/username/Library/Mobile Documents/iCloud\~md\~obsidian/Documents/Vault-New

旧 Vault:
 /Users/username/Library/Mobile Documents/iCloud\~md\~obsidian/Documents/Vault-Old
 
运行命令：

``` bash
go run main.go \
'/Users/username/Library/Mobile Documents/iCloud~md~obsidian/Documents/Vault-New' \
'/Users/username/Library/Mobile Documents/iCloud~md~obsidian/Documents/Vault-Old'
```
