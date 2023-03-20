CREATE TABLE IF NOT EXISTS reservation (
  reservation_id bigserial PRIMARY KEY,
  user_id        bigserial,
  reservation_date date,
  duration int,
  people_count int,
  approval BOOLEAN,
  notes text NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
