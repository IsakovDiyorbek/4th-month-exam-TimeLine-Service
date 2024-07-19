CREATE TABLE custom_events (        -- пользовательские события
    id UUID PRIMARY KEY,
    user_id UUID not NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    date DATE NOT NULL,
    category VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint default 0
);


CREATE TABLE milestones (   -- этап
    id UUID PRIMARY KEY,
    user_id UUID not NULL,
    title VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    category VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint default 0
);




HistoricalEvents kolleksiyasi:
{
  "_id": ObjectId(),
  "title": String,
  "date": Date,
  "category": String,
  "description": String,
  "source_url": String,
  "created_at": Date
}


UserTimeline kolleksiyasi:
{
  "_id": ObjectId(),
  "user_id": UUID,
  "events": [                                           --voqealar
    {
      "id": String,
      "type": String,
      "title": String,
      "date": Date,
      "preview": String                ---oldindan ko'rish
    }
  ],
  "last_updated": Date
}