
## グラフの日本語文字化け対応

# IPAexフォントのライセンス（IPAフォントライセンスv1.0）では、フォントファイルを 改変せずに 配布することが許可されている。
# ipaexg.ttf をプロジェクトのディレクトリに配置し、GitHub で公開して共有することは、ライセンスの範囲内で 問題ない。

# ただし、
# ライセンスファイルを含める: ipaexg.ttf と一緒に、ライセンスファイル (IPA_Font_License_Agreement_v1.0.txt) も必ずプロジェクトに含める。
# フォントファイルを変更しない: フォントファイル自体に変更を加えてはいけない。
# READMEにライセンス情報を記載: プロジェクトのREADMEファイルに、IPAexフォントを使用していること、およびライセンス情報を記載することを推奨される。


## グラフの調整

1. 図のサイズ:
figsize: plt.figure(figsize=(width, height)) で図のサイズをインチ単位で指定。
figsize=(10, 5) は幅10インチ、高さ5インチの図を作成。

2. フォントサイズ:
fontsize: plt.xlabel("日付", fontsize=14, fontproperties=jp_font) のように、fontsize 引数でフォントサイズを指定。
plt.ylabel、plt.title、plt.legend などでも同様に指定。
fontproperties で FontProperties オブジェクトを使う場合、jp_font = fm.FontProperties(fname=font_path, size=14) のように、size 引数でフォントサイズを指定。

3. 余白:
plt.subplots_adjust(): 図の余白を調整。
plt.subplots_adjust(left=0.1, bottom=0.1, right=0.9, top=0.9, wspace=0.2, hspace=0.2)
left, bottom, right, top: それぞれ左、下、右、上の余白を図の幅に対する割合で指定。
wspace, hspace: 複数 subplot を配置する場合、subplot 間の水平方向、垂直方向のスペースを図の幅/高さに対する割合で指定。
plt.tight_layout(): subplots_adjust より簡単に余白を自動調整。
ラベルなどが図の枠からはみ出さないように、自動的に余白を調整。

4. タイトルの位置:
plt.title("株価と指標", y=1.05, fontproperties=jp_font) のように、y 引数でタイトルの垂直方向の位置を調整。
1.0 より大きい値を指定すると、タイトルが図の上部に移動。

5. 凡例の位置とサイズ:
plt.legend(loc="upper left", fontsize=12, prop=jp_font) のように、loc 引数で凡例の位置を指定し、fontsize でフォントサイズを指定。
loc には、"upper left"、"lower right" などの方位を指定できます。 数値で位置を細かく指定する。
bbox_to_anchor と loc を組み合わせて、凡例の位置をより細かく調整することも可能。
plt.legend(bbox_to_anchor=(1.05, 1), loc='upper left', borderaxespad=0., fontsize=12, prop=jp_font)

plt.figure(figsize=(16, 8))で初めにサイズを決め、plt.plot() でグラフを描画した後、plt.savefig() を呼び出す前に、これらのオプションを設定。

