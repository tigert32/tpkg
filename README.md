#可以解包 抖音小程序 的tpkg文件

支持处理目录下的多个 `tpkg` 文件或单个文件，并且可以设定输出目录。

## 用法

```
Usage: -d 支持目录下多个pkg文件 -f 支持目录内单个文件 -o 设定输出目录
Usage: program -f sss.pkg | -d d:/appid_123/ver_123 [-o <output directory>]
```

### 参数说明

- `-d` : 支持目录下多个 `pkg` 文件处理。
- `-f` : 支持目录内单个文件处理。
- `-o` : 设定输出目录（可选）。

### 示例

处理单个文件：

```
program -f sss.pkg
```

处理目录下的多个文件，并设定输出目录：

```
program -d d:/ver_123 -o d:/output_directory
```

