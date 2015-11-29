package com.hifx.user.bloodcare.Routers;

public interface AsyncResponseInterface {

	void dataFromDB(String output, String functionname);
	void showToast();
	void getJson(String url, String code, String pagenumber, boolean saveToDB, String parentcode, String classfunctioname, String retainorDelete);
}