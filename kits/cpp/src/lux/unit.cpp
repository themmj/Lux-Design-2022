#include "lux/unit.hpp"

#include <cmath>

#include "lux/config.hpp"
#include "lux/observation.hpp"

namespace lux {
    int64_t Unit::moveCost(const Observation &obs, Direction direction) const {
        Position target = pos + Position::Delta(direction);
        if (target.x < 0 || target.y < 0 || static_cast<size_t>(target.x) >= obs.board.rubble.size()
            || static_cast<size_t>(target.y) >= obs.board.rubble.size()) {
            return -1;
        }
        auto factoryTeam = obs.board.factory_occupancy[target.y][target.x];
        if (factoryTeam != -1 && team_id != factoryTeam) {
            return -1;
        }
        auto rubble  = obs.board.rubble[target.y][target.x];
        auto weather = obs.getCurrentWeather();
        return std::ceil(
            (obs.config.ROBOTS[unit_type].MOVE_COST + obs.config.ROBOTS[unit_type].RUBBLE_MOVEMENT_COST * rubble)
            * weather.POWER_CONSUMPTION);
    }

    bool Unit::canMove(const Observation &obs, Direction direction) const {
        auto cost = moveCost(obs, direction);
        if (cost < 0) {
            return false;
        }
        return power >= cost;
    }

    int64_t Unit::digCost(const Observation &obs) const {
        auto weather = obs.getCurrentWeather();
        return std::ceil(obs.config.ROBOTS[unit_type].DIG_COST * weather.POWER_CONSUMPTION);
    }

    bool Unit::canDig(const Observation &obs) const { return power >= digCost(obs); }

    int64_t Unit::selfDestructCost(const Observation &obs) const {
        auto weather = obs.getCurrentWeather();
        return std::ceil(obs.config.ROBOTS[unit_type].SELF_DESTRUCT_COST * weather.POWER_CONSUMPTION);
    }

    bool Unit::canSelfDestruct(const Observation &obs) const { return power >= selfDestructCost(obs); }
}  // namespace lux
