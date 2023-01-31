export namespace services {
	
	export class EPW {
	    city?: string;
	    region?: string;
	    country?: string;
	    source?: string;
	    station_id?: string;
	    latitude?: number;
	    longitude?: number;
	    time_zone?: string;
	    elevation?: string;
	    year?: number[];
	    month?: number[];
	    day?: number[];
	    hour?: number[];
	    minute?: number[];
	    uncertainty?: string[];
	    dry_bulb_temperature?: number[];
	    dew_point_temperature?: number[];
	    relative_humidity?: number[];
	    atmospheric_station_pressure?: number[];
	    extraterrestrial_horizontal_radiation?: number[];
	    extraterrestrial_direct_normal_radiation?: number[];
	    horizontal_infrared_radiation_intensity?: number[];
	    global_horizontal_radiation?: number[];
	    direct_normal_radiation?: number[];
	    diffuse_horizontal_radiation?: number[];
	    global_horizontal_illuminance?: number[];
	    direct_normal_illuminance?: number[];
	    diffuse_horizontal_illuminance?: number[];
	    zenith_luminance?: number[];
	    wind_direction?: number[];
	    wind_speed?: number[];
	    total_sky_cover?: number[];
	    opaque_sky_cover?: number[];
	    visibility?: number[];
	    ceiling_height?: number[];
	    present_weather_observation?: number[];
	    present_weather_codes?: string[];
	    precipitable_water?: number[];
	    aerosol_optical_depth?: number[];
	    snow_depth?: number[];
	    days_since_last_snowfall?: number[];
	    albedo?: number[];
	    liquid_precipitation_depth?: number[];
	    liquid_precipitation_quantity?: number[];
	
	    static createFrom(source: any = {}) {
	        return new EPW(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.city = source["city"];
	        this.region = source["region"];
	        this.country = source["country"];
	        this.source = source["source"];
	        this.station_id = source["station_id"];
	        this.latitude = source["latitude"];
	        this.longitude = source["longitude"];
	        this.time_zone = source["time_zone"];
	        this.elevation = source["elevation"];
	        this.year = source["year"];
	        this.month = source["month"];
	        this.day = source["day"];
	        this.hour = source["hour"];
	        this.minute = source["minute"];
	        this.uncertainty = source["uncertainty"];
	        this.dry_bulb_temperature = source["dry_bulb_temperature"];
	        this.dew_point_temperature = source["dew_point_temperature"];
	        this.relative_humidity = source["relative_humidity"];
	        this.atmospheric_station_pressure = source["atmospheric_station_pressure"];
	        this.extraterrestrial_horizontal_radiation = source["extraterrestrial_horizontal_radiation"];
	        this.extraterrestrial_direct_normal_radiation = source["extraterrestrial_direct_normal_radiation"];
	        this.horizontal_infrared_radiation_intensity = source["horizontal_infrared_radiation_intensity"];
	        this.global_horizontal_radiation = source["global_horizontal_radiation"];
	        this.direct_normal_radiation = source["direct_normal_radiation"];
	        this.diffuse_horizontal_radiation = source["diffuse_horizontal_radiation"];
	        this.global_horizontal_illuminance = source["global_horizontal_illuminance"];
	        this.direct_normal_illuminance = source["direct_normal_illuminance"];
	        this.diffuse_horizontal_illuminance = source["diffuse_horizontal_illuminance"];
	        this.zenith_luminance = source["zenith_luminance"];
	        this.wind_direction = source["wind_direction"];
	        this.wind_speed = source["wind_speed"];
	        this.total_sky_cover = source["total_sky_cover"];
	        this.opaque_sky_cover = source["opaque_sky_cover"];
	        this.visibility = source["visibility"];
	        this.ceiling_height = source["ceiling_height"];
	        this.present_weather_observation = source["present_weather_observation"];
	        this.present_weather_codes = source["present_weather_codes"];
	        this.precipitable_water = source["precipitable_water"];
	        this.aerosol_optical_depth = source["aerosol_optical_depth"];
	        this.snow_depth = source["snow_depth"];
	        this.days_since_last_snowfall = source["days_since_last_snowfall"];
	        this.albedo = source["albedo"];
	        this.liquid_precipitation_depth = source["liquid_precipitation_depth"];
	        this.liquid_precipitation_quantity = source["liquid_precipitation_quantity"];
	    }
	}

}

