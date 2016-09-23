SCHEMA_FILES = user.capnp

schemas: $(SCHEMA_FILES)
	capnp compile -ogo:. $+
