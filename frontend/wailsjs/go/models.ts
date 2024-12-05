export namespace main {
	
	export class  {
	    // Go type: struct { DocumentID int "json:\"DOCUMENT_ID\""; DisplayDate int64 "json:\"DISPLAY_DATE\""; DocumentType string "json:\"DOCUMENT_TYPE\""; Grantors string "json:\"GRANTORS\""; Grantees string "json:\"GRANTEES\""; UnitNum interface {} "json:\"UNIT_NUM\"" }
	    attributes: any;
	
	    static createFrom(source: any = {}) {
	        return new (source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.attributes = this.convertValues(source["attributes"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class  {
	    name: string;
	    type: string;
	    alias: string;
	    sqlType: string;
	    domain: any;
	    defaultValue: any;
	    length?: number;
	
	    static createFrom(source: any = {}) {
	        return new (source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.type = source["type"];
	        this.alias = source["alias"];
	        this.sqlType = source["sqlType"];
	        this.domain = source["domain"];
	        this.defaultValue = source["defaultValue"];
	        this.length = source["length"];
	    }
	}
	export class AtlasDocuments {
	    objectIdFieldName: string;
	    // Go type: struct { Name string "json:\"name\""; IsSystemMaintained bool "json:\"isSystemMaintained\"" }
	    uniqueIdField: any;
	    globalIdFieldName: string;
	    geometryType: string;
	    // Go type: struct { Wkid int "json:\"wkid\""; LatestWkid int "json:\"latestWkid\"" }
	    spatialReference: any;
	    fields: [];
	    features: [];
	
	    static createFrom(source: any = {}) {
	        return new AtlasDocuments(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.objectIdFieldName = source["objectIdFieldName"];
	        this.uniqueIdField = this.convertValues(source["uniqueIdField"], Object);
	        this.globalIdFieldName = source["globalIdFieldName"];
	        this.geometryType = source["geometryType"];
	        this.spatialReference = this.convertValues(source["spatialReference"], Object);
	        this.fields = this.convertValues(source["fields"], );
	        this.features = this.convertValues(source["features"], );
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

