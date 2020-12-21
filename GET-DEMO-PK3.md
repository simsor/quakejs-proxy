# Download PK3 files from a QuakeJS server

This guide will walk you through getting the PK3 files from a QuakeJS content server.

## Prerequisites

- A Bash command-line (WSL works fine)
- The `jq` packages

## Steps

1. Download QuakeJS `get-assets.sh` script. It will be used to parse `manifest.json` and download every file in it.

```shell
$ wget https://raw.githubusercontent.com/begleysm/quakejs/master/html/get_assets.sh
```

2. Use this script to download assets from `content.quakejs.com` to the `./assets/` folder.

```shell
$ bash get_assets.sh
```

3. Run the `*-linuxq3ademo-1.11-6.x86.gz.sh` script to extract the Quake 3 Arena Linux demo (kill it when it hangs)

```shell
$ bash assets/*-linuxq3ademo-1.11-6.x86.gz.sh -target q3ademo
```

4. Do the same for the 1.32 update. Abort the installation when it asks for your root password, we don't want to actually install the update.

```shell
$ bash assets/*-linuxq3apoint-1.32b-3.x86.run --target pointrelease
```

5. Check the checksums

```shell
$ md5sum q3ademo/demoq3/pak0.pk3 pointrelease/baseq3/*.pk3
0613b3d4ef05e613a2b470571498690f  q3ademo/demoq3/pak0.pk3
48911719d91be25adb957f2d325db4a0  pointrelease/baseq3/pak1.pk3
d550ce896130c47166ca44b53f8a670a  pointrelease/baseq3/pak2.pk3
968dfd0f30dad67056115c8e92344ddc  pointrelease/baseq3/pak3.pk3
24bb1f4fcabd95f6e320c0e2f62f19ca  pointrelease/baseq3/pak4.pk3
734dcd06d2cbc7a16432ff6697f1c5ba  pointrelease/baseq3/pak5.pk3
873888a73055c023f6c38b8ca3f2ce05  pointrelease/baseq3/pak6.pk3
8fd38c53ed814b64f6ab03b5290965e4  pointrelease/baseq3/pak7.pk3
d8b96d429ca4a9c289071cb7e77e14d2  pointrelease/baseq3/pak8.pk3
```

6. Copy the PK3 files

```shell
$ cp q3ademo/demoq3/pak0.pk3 pointrelease/baseq3/*.pk3 /path/to/ioquake3/baseq3/
```

7. Optional: Also copy the files in `assets/baseq3/`, removing their prefix