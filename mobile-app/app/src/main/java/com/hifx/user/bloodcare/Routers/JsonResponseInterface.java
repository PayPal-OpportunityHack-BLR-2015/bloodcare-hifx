package com.hifx.user.bloodcare.Routers;

import org.json.JSONObject;

public interface JsonResponseInterface {

	void dataFromApi(JSONObject output, String classfunctionname);
	void ErrorResponseVolley(int errorcode, String functionname);;
	
}