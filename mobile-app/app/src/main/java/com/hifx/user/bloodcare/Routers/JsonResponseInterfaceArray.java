package com.hifx.user.bloodcare.Routers;

import org.json.JSONArray;

public interface JsonResponseInterfaceArray {

	void dataFromApi(JSONArray output, String classfunctionname);
	void ErrorResponseVolley(int errorcode, String functionname);;
	
}