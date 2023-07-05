FROM scratch

# バイナリをコピー
COPY kvs .

# 8080を使用する宣言
EXPOSE 8080

CMD ["/kvs"]