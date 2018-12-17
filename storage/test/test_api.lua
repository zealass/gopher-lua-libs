local storage = require("storage")
local inspect = require("inspect")
local time = require("time")

local s, err = storage.open("./test/db.json")
if err then error(err) end

-- set
local err = s:set("key", {"one", "two", 1}, 1)
if err then error(err) end

local err = s:set("key2", "value2", 60)
if err then error(err) end

local err = s:set("key3", 10.64, nil)
if err then error(err) end

local value, found, err = s:get("key")
if err then error(err) end
if not found then error("must be found") end

if not(value[1] == "one") then error("value") end
if not(value[3] == 1) then error("value") end

local value, found, err = s:get("key2")
if err then error(err) end
if not(value == "value2") then error("value") end

local value, found, err = s:get("key3")
if err then error(err) end
if not(value == 10.64) then error("value") end

time.sleep(1)

-- wait ttl
local value, found, err = s:get("key")
if err then error(err) end
if found then error("must be not found") end

-- close
local err = s:close()
if err then error(err) end