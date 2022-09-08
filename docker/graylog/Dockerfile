FROM graylog/graylog:4.3.5
COPY docker/graylog/content-pack-0ab6f1aa-cf2a-4b54-85c0-1f2c0735fa3e-1.json /usr/share/graylog/data/contentpacks/
ENV GRAYLOG_CONTENT_PACKS_AUTO_INSTALL content-pack-0ab6f1aa-cf2a-4b54-85c0-1f2c0735fa3e-1.json
ENV GRAYLOG_CONTENT_PACKS_LOADER_ENABLED true
ENV GRAYLOG_CONTENT_PACKS_DIR data/contentpacks