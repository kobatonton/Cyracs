# Cyracs
KVMを用いたVM自動構築システムです。

config ディレクトリ内に記述された定義ファイルをもとにVMの自動構築を行います。

ベースイメージは自身で作成することもできます。ただし、cloud-initが有効になっている必要があります。
<br><br><br>

現在対応しているOSは以下によってインターフェースが制御されているOSのみです。
- netplan
- network-scripts

# Installation
このツールを使用するにあたり、以下がインストールされている必要があります。
- libvirt v6.0.0 以上
- Go v1.16 以上


```
git clone https://github.com/kobatonton/cyracs.git
```

# Base images
https://drive.google.com/drive/folders/154yw9yI0jlLkY19tfbTyydiMZ2E26knq?usp=sharing

# Usage
## Create 
定義ファイルに書かれた仮想ネットワークと仮想マシン群を作成します。

```
go run main.go crate [filename]
```

### Example
```
#go run main.go create test.json


####Create Vnet#####
create vnet success: dmz
create vnet success: internal
create vnet success: work
####Create VM#####
create vm success: testvm
####Waiting start vms####
waiting start vm: testvm
start testvm successfully
```

## Destroy
定義ファイルに書かれた仮想ネットワークと仮想マシン群を削除します。

削除する際、仮想ネットワークと仮想マシンの強制終了が行われます。
```
go run main.go {crate,destroy,help} [filename]
```
### Example
```
#go run main.go destroy test.json

####Destroy VM#####
destroyed vm: testvm
####Destroy Vnet#####
destroyed vnet: dmz
destroyed vnet: internal
destroyed vnet: work
```

## Help
ヘルプを表示します。
```
go run main.go {crate,destroy,help} [filename]
```

