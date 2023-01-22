CREATE TABLE IF NOT EXISTS urlshortener (
  id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  short_url         VARCHAR(10) NOT NULL,
  long_url          text NOT NULL,
  number_of_clicks  INT DEFAULT 0,
  dt_created        TIMESTAMP NOT NULL DEFAULT now(),
  active            BOOLEAN DEFAULT TRUE
);

CREATE INDEX IF NOT EXISTS urlshortener_long_url ON urlshortener (long_url);
CREATE INDEX IF NOT EXISTS urlshortener_short_url ON urlshortener (short_url);
