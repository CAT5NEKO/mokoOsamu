# OSAMIZE

OSAMIZEは、指定されたパス内のファイルをサイズ順に精査し、大きい順に表示するツールです。ファイルの容量も表示します。

ある有名な格言、

### う、うわあああああああああああ（もこもこもこもこ）

から着想を得ました。

## 使用方法

```bash
./mokoOsamu.exe -osamize <intNum> <targetPath>
```

- intNum: 表示するファイル数を指定してください。
- targetPath: サイズを精査するディレクトリを指定してください。

### 現在抱えてる問題

SystemVolumeInformationというフォルダがあると、アクセス権限がないためエラーが発生します。そのため、このフォルダを無視するようにしています。

権限によりファイルの精査が拒否される場合は管理者権限で実行すると回避できます。

![image](https://github.com/CAT5NEKO/mokoOsamu/assets/111590457/5f40dd2f-5be5-46a9-9085-714e3fae7948)
