CREATE TABLE IF NOT EXISTS template_items (
    id serial PRIMARY KEY,
    template_id INTEGER NOT NULL REFERENCES templates(id) ON DELETE CASCADE,
    file_id INTEGER,
    organization_unit_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TRIGGER template_items_insert
AFTER INSERT ON template_items
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER template_items_update
AFTER UPDATE ON template_items
FOR EACH ROW
EXECUTE FUNCTION log_changes();

CREATE TRIGGER template_items_delete
AFTER DELETE ON template_items
FOR EACH ROW
EXECUTE FUNCTION log_changes();